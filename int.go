package cache

import "unsafe"

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: int
 * @Date:
 * @Desc: ...
 *
 */

type Int int

func (i Int) Len() int {
	size := unsafe.Sizeof(i)
	return int(size)
}
