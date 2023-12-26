package cache

/**
 * @Author: bbz
 * @Email: 2632141215@qq.com
 * @File: string
 * @Date:
 * @Desc: ...
 *
 */

type String string

func (d String) Len() int {
	return len(d)
}
