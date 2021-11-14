package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"runtime"
	"time"
)

// Level 预定义了日志等级
type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelError:
		return "error"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelPanic:
		return "panic"
	case LevelFatal:
		return "fatal"
	}

	return ""
}

// Logger 日志标准化
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithLevel(lvl Level) *Logger {
	ll := l.clone()
	ll.level = lvl
	return ll
}

// WithFields 设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

// WithContext 设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// WithCaller 设置当前某一层调用栈的信息（程序计数器、文件信息、行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}

	return ll
}

// WithCallersFrames 设置当前的整个调用栈的信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)

	//调用者用调用goroutine的堆栈上函数调用的返回程序计数器填充切片pc。
	//参数skip是在pc中记录之前要跳过的堆栈帧数，0表示调用者本身的帧，1表示调用者的调用者。
	//它返回写入pc的条目数。
	//要将这些pc转换成符号信息，如函数名和行号，请使用CallersFrames。
	//CallersFrames解释内联函数，并将返回的程序计数器调整为调用程序计数器。
	//直接在返回的pc上迭代是不鼓励的，就像在任何返回的pc上使用FuncForPC一样，
	//因为这些不能解释内联或返回程序计数器的调整。
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// 日志追踪
func (l *Logger) WithTrace() *Logger {
	ginCtx, ok := l.ctx.(*gin.Context)
	if ok {
		return l.WithFields(Fields{
			"trace_id": ginCtx.MustGet("X-Trace-ID"),
			"span_id":  ginCtx.MustGet("X-Span-ID"),
		})
	}
	return l
}

// JSONFormat 日志内容的格式化和日志输出动作的相关方法，继续写入如下代码
func (l *Logger) JSONFormat(message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; ok {
				data[k] = v
			}
		}
	}

	return data
}

func (l *Logger) Output(message string) {
	body, _ := json.Marshal(l.JSONFormat(message))
	content := string(body)
	switch l.level {
	case LevelDebug:
		l.newLogger.Printf(content)
	case LevelError:
		l.newLogger.Printf(content)
	case LevelInfo:
		l.newLogger.Printf(content)
	case LevelWarn:
		l.newLogger.Printf(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	}
}

//日志分级输出

func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelInfo).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelInfo).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelDebug).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelDebug).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}
func (l *Logger) Warn(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelWarn).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Warnf(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelWarn).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelError).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelError).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelPanic).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelPanic).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	l.WithLevel(LevelFatal).WithContext(ctx).WithTrace().Output(fmt.Sprint(v...))
}

func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	l.WithLevel(LevelFatal).WithContext(ctx).WithTrace().Output(fmt.Sprintf(format, v...))
}
