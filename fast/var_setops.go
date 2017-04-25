// -------------------------------------------------------------
// DO NOT EDIT! this file was generated automatically by gomacro
// Any change will be lost when the file is re-generated
// -------------------------------------------------------------

/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * var_setops.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/token"
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) varAddConst(upn int, index int, t r.Type, val I) {
	if isLiteralNumber(val, 0) || val == "" {
		return
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) += val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex128:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.ThreadGlobals.FileEnv.
							Binds[index]
						lhs.SetComplex(lhs.Complex() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case string:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.ThreadGlobals.FileEnv.
							Binds[index]
						lhs.SetString(lhs.String() +
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.ADD, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varAddExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) += fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.ThreadGlobals.FileEnv.
						Binds[index]
					lhs.SetComplex(lhs.Complex() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) string:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.ThreadGlobals.FileEnv.
						Binds[index]
					lhs.SetString(lhs.String() +
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.ADD, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varSubConst(upn int, index int, t r.Type, val I) {
	if isLiteralNumber(val, 0) {
		return
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) -= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex128:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.ThreadGlobals.FileEnv.
							Binds[index]
						lhs.SetComplex(lhs.Complex() -
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.SUB, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varSubExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) -= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.ThreadGlobals.FileEnv.
						Binds[index]
					lhs.SetComplex(lhs.Complex() -
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.SUB, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varMulConst(upn int, index int, t r.Type, val I) {
	if isLiteralNumber(val, 0) {

		c.varSetZero(upn, index, t)
		return
	} else if isLiteralNumber(val, 1) {
		return
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) *= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex128:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.ThreadGlobals.FileEnv.
							Binds[index]
						lhs.SetComplex(lhs.Complex() *
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.MUL, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varMulExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) *= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.ThreadGlobals.FileEnv.
						Binds[index]
					lhs.SetComplex(lhs.Complex() *
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.MUL, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varQuoConst(upn int, index int, t r.Type, val I) {
	if isLiteralNumber(val, 0) {
		c.Errorf("division by %v <%v>", val, t)
		return
	} else if isLiteralNumber(val, 1) {
		return
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			float64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) /= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			complex128:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.
							Outer.Outer.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					{
						lhs := o.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					{
						lhs := env.ThreadGlobals.FileEnv.
							Binds[index]
						lhs.SetComplex(lhs.Complex() /
							val,
						)
					}

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.QUO, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varQuoExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float32)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) float64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*float64)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*complex64)(unsafe.Pointer(&o.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) /= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) complex128:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.
						Outer.Outer.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				{
					lhs := o.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				{
					lhs := env.ThreadGlobals.FileEnv.
						Binds[index]
					lhs.SetComplex(lhs.Complex() /
						fun(env),
					)
				}

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.QUO, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varRemConst(upn int, index int, t r.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, 0) {
			c.Errorf("division by %v <%v>", val, t)
			return
		} else if isLiteralNumber(val, 1) {

			c.varSetZero(upn, index, t)
			return
		}
	}
	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) %= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.REM, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varRemExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) %= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.REM, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varAndConst(upn int, index int, t r.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {
			return
		} else if isLiteralNumber(val, 0) {

			c.varSetZero(upn, index, t)
			return
		}
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varAndExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varOrConst(upn int, index int, t r.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		return
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) |= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.OR, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varOrExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) |= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.OR, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varXorConst(upn int, index int, t r.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) && isLiteralNumber(val, 0) {
		return
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) ^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.XOR, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varXorExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) ^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.XOR, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) varAndnotConst(upn int, index int, t r.Type, val I) {
	if IsCategory(t.Kind(), r.Int, r.Uint) {
		if isLiteralNumber(val, -1) {

			c.varSetZero(upn, index, t)
			return
		} else if isLiteralNumber(val, 0) {
			return
		}
	}

	{
		var ret Stmt
		switch val := val.(type) {
		case int:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case int64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint8:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint16:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint32:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uint64:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case

			uintptr:
			switch upn {
			case 0:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case 2:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			default:

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			case c.Depth - 1:

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) &^= val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		default:
			c.Errorf(`invalid operator %s= on <%v>`, token.AND_NOT, t)

		}
		c.Code.Append(ret)
	}
}
func (c *Comp) varAndnotExpr(upn int, index int, t r.Type, fun I) {
	var ret Stmt
	switch fun := fun.(type) {
	case func(*Env) int:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int8)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int16)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int32)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) int64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*int64)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint8:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint8)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint16:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint16)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint32:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uint32)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uint64:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					IntBinds[index] &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case func(*Env) uintptr:
		switch upn {
		case 0:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case 2:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.
					Outer.Outer.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				*(*uintptr)(unsafe.Pointer(&o.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		case c.Depth - 1:

			ret = func(env *Env) (Stmt, *Env) {
				*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
					IntBinds[index])) &^= fun(env)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		c.Errorf(`invalid operator %s= between <%v> and <%v>`, token.AND_NOT, t, funTypeOuts(fun))

	}
	c.Code.Append(ret)
}
func (c *Comp) SetVar(va *Var, op token.Token, init *Expr) {
	t := va.Type
	if init.Const() {
		init.ConstTo(t)
	} else if !init.Type.AssignableTo(t) {
		c.Errorf("incompatible types in assignment: <%v> %s <%v>", t, op, init.Type)
		return
	}
	class := va.Desc.Class()
	if class != VarBind && class != IntBind {
		c.Errorf("invalid operator %s on %v", op, class)
		return
	}
	upn := va.Upn
	index := va.Desc.Index()
	if index == NoIndex {
		if op != token.ASSIGN {
			c.Errorf("invalid operator %s on _", op)
		}

		if !init.Const() {
			c.Code.Append(init.AsStmt())
		}

		return
	}
	if init.Const() {
		val := init.Value
		v := r.ValueOf(val)
		if v == None || v == Nil {
			v = r.Zero(t)
			val = v.Interface()
		} else if v.Type() != t {
			v = v.Convert(t)
			val = v.Interface()
		}
		switch op {
		case token.ASSIGN:
			c.varSetConst(upn, index, t, val)
		case token.ADD, token.ADD_ASSIGN:
			c.varAddConst(upn, index, t, val)
		case token.SUB, token.SUB_ASSIGN:
			c.varSubConst(upn, index, t, val)
		case token.MUL, token.MUL_ASSIGN:
			c.varMulConst(upn, index, t, val)
		case token.QUO, token.QUO_ASSIGN:
			c.varQuoConst(upn, index, t, val)
		case token.REM, token.REM_ASSIGN:
			c.varRemConst(upn, index, t, val)
		case token.AND, token.AND_ASSIGN:
			c.varAndConst(upn, index, t, val)
		case token.OR, token.OR_ASSIGN:
			c.varOrConst(upn, index, t, val)
		case token.XOR, token.XOR_ASSIGN:
			c.varAndConst(upn, index, t, val)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.varAndnotConst(upn, index, t, val)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	} else {
		fun := init.Fun
		switch op {
		case token.ASSIGN:
			c.varSetExpr(upn, index, t, fun)
		case token.ADD, token.ADD_ASSIGN:
			c.varAddExpr(upn, index, t, fun)
		case token.SUB, token.SUB_ASSIGN:
			c.varSubExpr(upn, index, t, fun)
		case token.MUL, token.MUL_ASSIGN:
			c.varMulExpr(upn, index, t, fun)
		case token.QUO, token.QUO_ASSIGN:
			c.varQuoExpr(upn, index, t, fun)
		case token.REM, token.REM_ASSIGN:
			c.varRemExpr(upn, index, t, fun)
		case token.AND, token.AND_ASSIGN:
			c.varAndExpr(upn, index, t, fun)
		case token.OR, token.OR_ASSIGN:
			c.varOrExpr(upn, index, t, fun)
		case token.XOR, token.XOR_ASSIGN:
			c.varAndExpr(upn, index, t, fun)
		case token.AND_NOT, token.AND_NOT_ASSIGN:
			c.varAndnotExpr(upn, index, t, fun)
		default:
			c.Errorf("operator %s is not implemented", op)
		}
	}
}