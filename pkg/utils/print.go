// Copyright © 2022 zc2638 <zc2638@qq.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"context"
	"io"
	"os"
	"time"
)

func VerbatimPrint(ctx context.Context, s string, delay time.Duration) {
	delay = delay * time.Millisecond
	for _, v := range s {
		_ = writeString(os.Stdout, string(v))

		select {
		case <-ctx.Done():
			return
		case <-time.After(delay):
		}
	}
}

func ClearPrint() error {
	return writeString(os.Stdout, "\033[2K\r")
}

func writeString(writer io.Writer, str string) error {
	if _, err := io.WriteString(writer, str); err != nil {
		return err
	}

	if f, ok := writer.(*os.File); ok {
		// ignore any errors in Sync(), as stdout
		// can't be synced on some operating systems
		// like Debian 9 (Stretch)
		f.Sync()
	}
	return nil
}
