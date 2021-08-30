package seal

import "io"

//Unseal read first byte , wich mean a lenth of data , its need to be because if we want to use a ReadAll to read data , it blocking forever
func Unseal(conn io.Reader) ([]byte, error) {

	buf := make([]byte, 1)

	_, err := conn.Read(buf)

	if err != nil && err != io.EOF {
		return []byte{}, err
	}

	data := make([]byte, int(buf[0]))

	_, err = conn.Read(data)

	if err != nil && err != io.EOF {
		return []byte{}, err
	}

	return data, nil
}

//Seal add first byte , wich mean a lenth of data , its need to be because if we want to use a ReadAll to read data , it blocking forever
func Seal(json []byte) []byte {
	n := len(json)
	return append([]byte{byte(n)}, json...)
}
