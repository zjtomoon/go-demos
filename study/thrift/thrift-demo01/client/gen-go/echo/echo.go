// Code generated by Thrift Compiler (0.18.1). DO NOT EDIT.

package echo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"regexp"
	"strings"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = errors.New
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// (needed by validator.)
var _ = strings.Contains
var _ = regexp.MatchString

// Attributes:
//   - Msg
type EchoReq struct {
	Msg string `thrift:"msg,1" db:"msg" json:"msg"`
}

func NewEchoReq() *EchoReq {
	return &EchoReq{}
}

func (p *EchoReq) GetMsg() string {
	return p.Msg
}
func (p *EchoReq) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoReq) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *EchoReq) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EchoReq"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoReq) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "msg", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:msg: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Msg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.msg (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:msg: ", p), err)
	}
	return err
}

func (p *EchoReq) Equals(other *EchoReq) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Msg != other.Msg {
		return false
	}
	return true
}

func (p *EchoReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoReq(%+v)", *p)
}

func (p *EchoReq) Validate() error {
	return nil
}

// Attributes:
//   - Msg
type EchoRes struct {
	Msg string `thrift:"msg,1" db:"msg" json:"msg"`
}

func NewEchoRes() *EchoRes {
	return &EchoRes{}
}

func (p *EchoRes) GetMsg() string {
	return p.Msg
}
func (p *EchoRes) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoRes) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Msg = v
	}
	return nil
}

func (p *EchoRes) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "EchoRes"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoRes) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "msg", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:msg: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Msg)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.msg (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:msg: ", p), err)
	}
	return err
}

func (p *EchoRes) Equals(other *EchoRes) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Msg != other.Msg {
		return false
	}
	return true
}

func (p *EchoRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoRes(%+v)", *p)
}

func (p *EchoRes) Validate() error {
	return nil
}

// Attributes:
//   - ID
type Num struct {
	ID int32 `thrift:"id,1,required" db:"id" json:"id"`
}

func NewNum() *Num {
	return &Num{}
}

func (p *Num) GetID() int32 {
	return p.ID
}
func (p *Num) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetID bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
				issetID = true
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetID {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field ID is not set"))
	}
	return nil
}

func (p *Num) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ID = v
	}
	return nil
}

func (p *Num) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Num"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Num) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "id", thrift.I32, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:id: ", p), err)
	}
	if err := oprot.WriteI32(ctx, int32(p.ID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.id (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:id: ", p), err)
	}
	return err
}

func (p *Num) Equals(other *Num) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.ID != other.ID {
		return false
	}
	return true
}

func (p *Num) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Num(%+v)", *p)
}

func (p *Num) Validate() error {
	return nil
}

type Echo interface {
	// Parameters:
	//  - Req
	Echo(ctx context.Context, req *EchoReq) (_r *EchoRes, _err error)
	// Parameters:
	//  - Num1
	//  - Num2
	Add(ctx context.Context, num1 *Num, num2 *Num) (_r *Num, _err error)
}

type EchoClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewEchoClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *EchoClient {
	return &EchoClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewEchoClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *EchoClient {
	return &EchoClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewEchoClient(c thrift.TClient) *EchoClient {
	return &EchoClient{
		c: c,
	}
}

func (p *EchoClient) Client_() thrift.TClient {
	return p.c
}

func (p *EchoClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *EchoClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//   - Req
func (p *EchoClient) Echo(ctx context.Context, req *EchoReq) (_r *EchoRes, _err error) {
	var _args0 EchoEchoArgs
	_args0.Req = req
	var _result2 EchoEchoResult
	var _meta1 thrift.ResponseMeta
	_meta1, _err = p.Client_().Call(ctx, "echo", &_args0, &_result2)
	p.SetLastResponseMeta_(_meta1)
	if _err != nil {
		return
	}
	if _ret3 := _result2.GetSuccess(); _ret3 != nil {
		return _ret3, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "echo failed: unknown result")
}

// Parameters:
//   - Num1
//   - Num2
func (p *EchoClient) Add(ctx context.Context, num1 *Num, num2 *Num) (_r *Num, _err error) {
	var _args4 EchoAddArgs
	_args4.Num1 = num1
	_args4.Num2 = num2
	var _result6 EchoAddResult
	var _meta5 thrift.ResponseMeta
	_meta5, _err = p.Client_().Call(ctx, "Add", &_args4, &_result6)
	p.SetLastResponseMeta_(_meta5)
	if _err != nil {
		return
	}
	if _ret7 := _result6.GetSuccess(); _ret7 != nil {
		return _ret7, nil
	}
	return nil, thrift.NewTApplicationException(thrift.MISSING_RESULT, "Add failed: unknown result")
}

type EchoProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Echo
}

func (p *EchoProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *EchoProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *EchoProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewEchoProcessor(handler Echo) *EchoProcessor {

	self8 := &EchoProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self8.processorMap["echo"] = &echoProcessorEcho{handler: handler}
	self8.processorMap["Add"] = &echoProcessorAdd{handler: handler}
	return self8
}

func (p *EchoProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x9 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x9.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x9

}

type echoProcessorEcho struct {
	handler Echo
}

func (p *echoProcessorEcho) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	var _write_err10 error
	args := EchoEchoArgs{}
	if err2 := args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "echo", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := EchoEchoResult{}
	if retval, err2 := p.handler.Echo(ctx, args.Req); err2 != nil {
		tickerCancel()
		err = thrift.WrapTException(err2)
		if errors.Is(err2, thrift.ErrAbandonRequest) {
			return false, thrift.WrapTException(err2)
		}
		_exc11 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing echo: "+err2.Error())
		if err2 := oprot.WriteMessageBegin(ctx, "echo", thrift.EXCEPTION, seqId); err2 != nil {
			_write_err10 = thrift.WrapTException(err2)
		}
		if err2 := _exc11.Write(ctx, oprot); _write_err10 == nil && err2 != nil {
			_write_err10 = thrift.WrapTException(err2)
		}
		if err2 := oprot.WriteMessageEnd(ctx); _write_err10 == nil && err2 != nil {
			_write_err10 = thrift.WrapTException(err2)
		}
		if err2 := oprot.Flush(ctx); _write_err10 == nil && err2 != nil {
			_write_err10 = thrift.WrapTException(err2)
		}
		if _write_err10 != nil {
			return false, thrift.WrapTException(_write_err10)
		}
		return true, err
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 := oprot.WriteMessageBegin(ctx, "echo", thrift.REPLY, seqId); err2 != nil {
		_write_err10 = thrift.WrapTException(err2)
	}
	if err2 := result.Write(ctx, oprot); _write_err10 == nil && err2 != nil {
		_write_err10 = thrift.WrapTException(err2)
	}
	if err2 := oprot.WriteMessageEnd(ctx); _write_err10 == nil && err2 != nil {
		_write_err10 = thrift.WrapTException(err2)
	}
	if err2 := oprot.Flush(ctx); _write_err10 == nil && err2 != nil {
		_write_err10 = thrift.WrapTException(err2)
	}
	if _write_err10 != nil {
		return false, thrift.WrapTException(_write_err10)
	}
	return true, err
}

type echoProcessorAdd struct {
	handler Echo
}

func (p *echoProcessorAdd) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	var _write_err12 error
	args := EchoAddArgs{}
	if err2 := args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "Add", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := EchoAddResult{}
	if retval, err2 := p.handler.Add(ctx, args.Num1, args.Num2); err2 != nil {
		tickerCancel()
		err = thrift.WrapTException(err2)
		if errors.Is(err2, thrift.ErrAbandonRequest) {
			return false, thrift.WrapTException(err2)
		}
		_exc13 := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Add: "+err2.Error())
		if err2 := oprot.WriteMessageBegin(ctx, "Add", thrift.EXCEPTION, seqId); err2 != nil {
			_write_err12 = thrift.WrapTException(err2)
		}
		if err2 := _exc13.Write(ctx, oprot); _write_err12 == nil && err2 != nil {
			_write_err12 = thrift.WrapTException(err2)
		}
		if err2 := oprot.WriteMessageEnd(ctx); _write_err12 == nil && err2 != nil {
			_write_err12 = thrift.WrapTException(err2)
		}
		if err2 := oprot.Flush(ctx); _write_err12 == nil && err2 != nil {
			_write_err12 = thrift.WrapTException(err2)
		}
		if _write_err12 != nil {
			return false, thrift.WrapTException(_write_err12)
		}
		return true, err
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 := oprot.WriteMessageBegin(ctx, "Add", thrift.REPLY, seqId); err2 != nil {
		_write_err12 = thrift.WrapTException(err2)
	}
	if err2 := result.Write(ctx, oprot); _write_err12 == nil && err2 != nil {
		_write_err12 = thrift.WrapTException(err2)
	}
	if err2 := oprot.WriteMessageEnd(ctx); _write_err12 == nil && err2 != nil {
		_write_err12 = thrift.WrapTException(err2)
	}
	if err2 := oprot.Flush(ctx); _write_err12 == nil && err2 != nil {
		_write_err12 = thrift.WrapTException(err2)
	}
	if _write_err12 != nil {
		return false, thrift.WrapTException(_write_err12)
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//   - Req
type EchoEchoArgs struct {
	Req *EchoReq `thrift:"req,1" db:"req" json:"req"`
}

func NewEchoEchoArgs() *EchoEchoArgs {
	return &EchoEchoArgs{}
}

var EchoEchoArgs_Req_DEFAULT *EchoReq

func (p *EchoEchoArgs) GetReq() *EchoReq {
	if !p.IsSetReq() {
		return EchoEchoArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *EchoEchoArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *EchoEchoArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoEchoArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Req = &EchoReq{}
	if err := p.Req.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Req), err)
	}
	return nil
}

func (p *EchoEchoArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "echo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoEchoArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "req", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:req: ", p), err)
	}
	if err := p.Req.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Req), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:req: ", p), err)
	}
	return err
}

func (p *EchoEchoArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoEchoArgs(%+v)", *p)
}

// Attributes:
//   - Success
type EchoEchoResult struct {
	Success *EchoRes `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewEchoEchoResult() *EchoEchoResult {
	return &EchoEchoResult{}
}

var EchoEchoResult_Success_DEFAULT *EchoRes

func (p *EchoEchoResult) GetSuccess() *EchoRes {
	if !p.IsSetSuccess() {
		return EchoEchoResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoEchoResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoEchoResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoEchoResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &EchoRes{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *EchoEchoResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "echo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoEchoResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *EchoEchoResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoEchoResult(%+v)", *p)
}

// Attributes:
//   - Num1
//   - Num2
type EchoAddArgs struct {
	Num1 *Num `thrift:"num1,1" db:"num1" json:"num1"`
	Num2 *Num `thrift:"num2,2" db:"num2" json:"num2"`
}

func NewEchoAddArgs() *EchoAddArgs {
	return &EchoAddArgs{}
}

var EchoAddArgs_Num1_DEFAULT *Num

func (p *EchoAddArgs) GetNum1() *Num {
	if !p.IsSetNum1() {
		return EchoAddArgs_Num1_DEFAULT
	}
	return p.Num1
}

var EchoAddArgs_Num2_DEFAULT *Num

func (p *EchoAddArgs) GetNum2() *Num {
	if !p.IsSetNum2() {
		return EchoAddArgs_Num2_DEFAULT
	}
	return p.Num2
}
func (p *EchoAddArgs) IsSetNum1() bool {
	return p.Num1 != nil
}

func (p *EchoAddArgs) IsSetNum2() bool {
	return p.Num2 != nil
}

func (p *EchoAddArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoAddArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	p.Num1 = &Num{}
	if err := p.Num1.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Num1), err)
	}
	return nil
}

func (p *EchoAddArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	p.Num2 = &Num{}
	if err := p.Num2.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Num2), err)
	}
	return nil
}

func (p *EchoAddArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Add_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoAddArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "num1", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:num1: ", p), err)
	}
	if err := p.Num1.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Num1), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:num1: ", p), err)
	}
	return err
}

func (p *EchoAddArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "num2", thrift.STRUCT, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:num2: ", p), err)
	}
	if err := p.Num2.Write(ctx, oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Num2), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:num2: ", p), err)
	}
	return err
}

func (p *EchoAddArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoAddArgs(%+v)", *p)
}

// Attributes:
//   - Success
type EchoAddResult struct {
	Success *Num `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewEchoAddResult() *EchoAddResult {
	return &EchoAddResult{}
}

var EchoAddResult_Success_DEFAULT *Num

func (p *EchoAddResult) GetSuccess() *Num {
	if !p.IsSetSuccess() {
		return EchoAddResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoAddResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoAddResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *EchoAddResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &Num{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *EchoAddResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Add_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *EchoAddResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *EchoAddResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoAddResult(%+v)", *p)
}
