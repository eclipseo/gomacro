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
 * var_set.go
 *
 *  Created on Apr 09, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	r "reflect"
	"unsafe"

	. "github.com/cosmos72/gomacro/base"
)

func (c *Comp) varSetZero(upn int, index int, t r.Type) {
	zero := r.Zero(t).Interface()
	c.varSetConst(upn, index, t, zero)
}
func (c *Comp) varSetConst(upn int, index int, t r.Type, val I) {
	v := r.ValueOf(val)
	if ValueType(v) == nil {
		v = r.Zero(t)
	} else {
		v = v.Convert(t)
	}

	var ret func(env *Env) (Stmt, *Env)
	switch upn {
	case 0:
		switch t {
		case TypeOfBool:

			{
				val := v.Interface().(bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:
			{
				val := v.Interface().(int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:
			{
				val := v.Interface().(int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:
			{
				val := v.Interface().(int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				val := v.Interface().(int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				val := v.Interface().(int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				val := v.Interface().(uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint8:
			{
				val := v.Interface().(uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint16:
			{
				val := v.Interface().(uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint32:
			{
				val := v.Interface().(uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				val := v.Interface().(uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				val := v.Interface().(uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				val := v.Interface().(float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				val := v.Interface().(float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				val := v.Interface().(complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				val := v.Interface().(complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				val := v.Interface().(string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case 1:
		switch t {
		case TypeOfBool:

			{
				val := v.Interface().(bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:
			{
				val := v.Interface().(int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:
			{
				val := v.Interface().(int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:
			{
				val := v.Interface().(int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				val := v.Interface().(int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				val := v.Interface().(int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				val := v.Interface().(uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint8:
			{
				val := v.Interface().(uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint16:
			{
				val := v.Interface().(uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint32:
			{
				val := v.Interface().(uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				val := v.Interface().(uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				val := v.Interface().(uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				val := v.Interface().(float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				val := v.Interface().(float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				val := v.Interface().(complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				val := v.Interface().(complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				val := v.Interface().(string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case 2:
		switch t {
		case TypeOfBool:

			{
				val := v.Interface().(bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:
			{
				val := v.Interface().(int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:
			{
				val := v.Interface().(int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:
			{
				val := v.Interface().(int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				val := v.Interface().(int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				val := v.Interface().(int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				val := v.Interface().(uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint8:
			{
				val := v.Interface().(uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint16:
			{
				val := v.Interface().(uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint32:
			{
				val := v.Interface().(uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				val := v.Interface().(uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				val := v.Interface().(uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				val := v.Interface().(float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				val := v.Interface().(float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				val := v.Interface().(complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				val := v.Interface().(complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				val := v.Interface().(string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.
					Outer.Outer.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	default:
		switch t {
		case TypeOfBool:

			{
				val := v.Interface().(bool)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*bool)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:
			{
				val := v.Interface().(int)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:
			{
				val := v.Interface().(int8)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:
			{
				val := v.Interface().(int16)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				val := v.Interface().(int32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				val := v.Interface().(int64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				val := v.Interface().(uint)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint8:
			{
				val := v.Interface().(uint8)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint16:
			{
				val := v.Interface().(uint16)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint32:
			{
				val := v.Interface().(uint32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				val := v.Interface().(uint64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				val := v.Interface().(uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				val := v.Interface().(float32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				val := v.Interface().(float64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				val := v.Interface().(complex64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				val := v.Interface().(complex128)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				val := v.Interface().(string)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				o := env.Outer.Outer.Outer
				for i := 3; i < upn; i++ {
					o = o.Outer
				}

				o.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	case c.Depth - 1:
		switch t {
		case TypeOfBool:

			{
				val := v.Interface().(bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:
			{
				val := v.Interface().(int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:
			{
				val := v.Interface().(int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:
			{
				val := v.Interface().(int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				val := v.Interface().(int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				val := v.Interface().(int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				val := v.Interface().(uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint8:
			{
				val := v.Interface().(uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint16:
			{
				val := v.Interface().(uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint32:
			{
				val := v.Interface().(uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				val := v.Interface().(uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				val := v.Interface().(uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				val := v.Interface().(float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				val := v.Interface().(float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				val := v.Interface().(complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = val

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				val := v.Interface().(complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetComplex(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				val := v.Interface().(string)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetString(val,
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:

			ret = func(env *Env) (Stmt, *Env) {
				env.ThreadGlobals.FileEnv.
					Binds[index].Set(v,
				)

				env.IP++
				return env.Code[env.IP], env
			}
		}
	}
	c.Code.Append(ret)
}
func (c *Comp) varSetExpr(upn int, index int, t r.Type, fun I) {
	var ret func(env *Env) (Stmt, *Env)
	switch upn {
	case 0:
		switch t {
		case TypeOfBool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := funAsX1(fun)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Binds[index].Set(fun(env).Convert(t),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	case 1:
		switch t {
		case TypeOfBool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := funAsX1(fun)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.
						Binds[index].Set(fun(env).Convert(t),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	case 2:
		switch t {
		case TypeOfBool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.
						Outer.Outer.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := funAsX1(fun)

				ret = func(env *Env) (Stmt, *Env) {
					env.
						Outer.Outer.
						Binds[index].Set(fun(env).Convert(t),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	default:
		switch t {
		case TypeOfBool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*bool)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int8)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int16)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int32)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*int64)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint8)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint16)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uint32)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*uintptr)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float32)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*float64)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					*(*complex64)(unsafe.Pointer(&o.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := funAsX1(fun)

				ret = func(env *Env) (Stmt, *Env) {
					o := env.Outer.Outer.Outer
					for i := 3; i < upn; i++ {
						o = o.Outer
					}

					o.
						Binds[index].Set(fun(env).Convert(t),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	case c.Depth - 1:
		switch t {
		case TypeOfBool:

			{
				fun := fun.(func(*Env) bool)

				ret = func(env *Env) (Stmt, *Env) {
					*(*bool)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt:

			{
				fun := fun.(func(*Env) int)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt8:

			{
				fun := fun.(func(*Env) int8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt16:

			{
				fun := fun.(func(*Env) int16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt32:
			{
				fun := fun.(func(*Env) int32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfInt64:
			{
				fun := fun.(func(*Env) int64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*int64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint:
			{
				fun := fun.(func(*Env) uint)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint8:
			{
				fun := fun.(func(*Env) uint8)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint8)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint16:
			{
				fun := fun.(func(*Env) uint16)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint16)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}
		case TypeOfUint32:
			{
				fun := fun.(func(*Env) uint32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uint32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUint64:
			{
				fun := fun.(func(*Env) uint64)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						IntBinds[index] = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfUintptr:
			{
				fun := fun.(func(*Env) uintptr)

				ret = func(env *Env) (Stmt, *Env) {
					*(*uintptr)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat32:
			{
				fun := fun.(func(*Env) float32)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float32)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfFloat64:
			{
				fun := fun.(func(*Env) float64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*float64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex64:
			{
				fun := fun.(func(*Env) complex64)

				ret = func(env *Env) (Stmt, *Env) {
					*(*complex64)(unsafe.Pointer(&env.ThreadGlobals.FileEnv.
						IntBinds[index])) = fun(env)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfComplex128:
			{
				fun := fun.(func(*Env) complex128)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetComplex(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		case TypeOfString:
			{
				fun := fun.(func(*Env) string)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].SetString(fun(env),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		default:
			{
				fun := funAsX1(fun)

				ret = func(env *Env) (Stmt, *Env) {
					env.ThreadGlobals.FileEnv.
						Binds[index].Set(fun(env).Convert(t),
					)

					env.IP++
					return env.Code[env.IP], env
				}
			}

		}
	}
	c.Code.Append(ret)
}