package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"../Entity"
)

type Transfer struct {
	Conn net.Conn
}

func (this *Transfer) ReadMsg() (msg *Entity.Message, err error) {
	msg = new(Entity.Message)
	length_buf := make([]byte, 4)
	_, err = this.Conn.Read(length_buf)
	if err != nil {
		fmt.Println(err)
		// sendErr(conn, err)
		return nil, err
	}
	length := binary.BigEndian.Uint32(length_buf)
	data_buf := make([]byte, length)
	_, err = this.Conn.Read(data_buf)
	if err != nil {
		fmt.Println(err)
		// sendErr(conn, err)
		return nil, err
	}
	err = json.Unmarshal(data_buf, msg)
	if err != nil {
		fmt.Println(err)
		// sendErr(conn, err)
		return nil, err
	}
	return
}
func (this *Transfer) SendMsg(msg *Entity.Message) (err error) {
	msg_b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	length := uint32(len(msg_b))
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, length)
	_, err = this.Conn.Write(buf)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(msg_b)
	if err != nil {
		fmt.Println("request error", err)
		return err
	}
	fmt.Println(length)
	return nil
}
func (this *Transfer) WriteResult(result *Entity.ResultType) (err error) {
	result_buf, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		// sendErr(conn, err)
		return err
	}
	result_length := uint32(len(result_buf))
	length_buf := make([]byte, 4)
	binary.BigEndian.PutUint32(length_buf, result_length)
	_, err = this.Conn.Write(length_buf)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(result_buf)
	if err != nil {
		fmt.Println("request error", err)
		return err
	}
	return nil
}
