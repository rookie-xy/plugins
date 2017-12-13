package console

import "bufio"

func client(writer *bufio.Writer, buffer []byte, end byte) error {
	written := 0
	for written < len(buffer) {
		n, err := writer.Write(buffer[written:])
		if err != nil {
			return err
		}

		written += n
	}

	if err := writer.WriteByte(end); err != nil {
		return err
	}

	return nil
}
