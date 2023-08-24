package slice

import "errors"

// 第一第三条实现了 第四条要求缩容暂时没实现
func Pop[T any](S []T, idx int) ([]T, error) {
	l := len(S)
	if idx >= l {
		return S, errors.New("传入的下标必须小于切片的长度")
	}
	res := append(S[:idx], S[idx+1:]...)
	return res, nil
}
