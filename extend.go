package yaml

//// DecodeF does the same thing as Decode, but it does not consider if the v is an instance of Unmarshaler
//// and enforces default decoding method.
//func (n *Node) DecodeF(v interface{}) (err error) {
//	d := newDecoder()
//	defer handleErr(&err)
//	out := reflect.ValueOf(v)
//	if out.Kind() == reflect.Ptr && !out.IsNil() {
//		out = out.Elem()
//	}
//	d.unmarshalF(n, out)
//	if len(d.terrors) > 0 {
//		return &TypeError{d.terrors}
//	}
//	return nil
//}
//
//// unmarshalF does the same thing as d.unmarshal, except that is calls `d.prepareF()` instead of `d.prepare()`
//func (d *decoder) unmarshalF(n *Node, out reflect.Value) (good bool) {
//	d.decodeCount++
//	if d.aliasDepth > 0 {
//		d.aliasCount++
//	}
//	if d.aliasCount > 100 && d.decodeCount > 1000 && float64(d.aliasCount)/float64(d.decodeCount) > allowedAliasRatio(d.decodeCount) {
//		failf("document contains excessive aliasing")
//	}
//	if out.Type() == nodeType {
//		out.Set(reflect.ValueOf(n).Elem())
//		return true
//	}
//	switch n.Kind {
//	case DocumentNode:
//		return d.document(n, out)
//	case AliasNode:
//		return d.alias(n, out)
//	}
//	out, unmarshaled, good := d.prepareF(n, out)
//	if unmarshaled {
//		return good
//	}
//	switch n.Kind {
//	case ScalarNode:
//		good = d.scalar(n, out)
//	case MappingNode:
//		good = d.mapping(n, out)
//	case SequenceNode:
//		good = d.sequence(n, out)
//	case 0:
//		if n.IsZero() {
//			return d.null(out)
//		}
//		fallthrough
//	default:
//		failf("cannot decode node with unknown kind %d", n.Kind)
//	}
//	return good
//}
//
//// prepareF does the same thing as `d.prepare`, but does not check for the possible Unmarshaler type values, and
//// always returns `unmarshaled == false`, effectively enforcing standard unmarshalling at the level above.
//func (d *decoder) prepareF(n *Node, out reflect.Value) (newout reflect.Value, unmarshaled, good bool) {
//	if n.ShortTag() == nullTag {
//		return out, false, false
//	}
//	again := true
//	for again {
//		again = false
//		if out.Kind() == reflect.Ptr {
//			if out.IsNil() {
//				out.Set(reflect.New(out.Type().Elem()))
//			}
//			out = out.Elem()
//			again = true
//		}
//	}
//	return out, false, false
//}
