package fp

func Curry2[T any, K any, R any](fn func(T, K) R) func(K) func(T) R {
	return func(k K) func(T) R {
		return func(t T) R {
			return fn(t, k)
		}
	}
}

func Curry3[T any, K any, A any, R any](fn func(T, K, A) R) func(A) func(K) func(T) R {
	return func(a A) func(K) func(T) R {
		return func(k K) func(T) R {
			return func(t T) R {
				return fn(t, k, a)
			}
		}
	}
}

func Curry4[T any, K any, A any, B any, R any](fn func(T, K, A, B) R) func(B) func(A) func(K) func(T) R {
	return func(b B) func(A) func(K) func(T) R {
		return func(a A) func(K) func(T) R {
			return func(k K) func(T) R {
				return func(t T) R {
					return fn(t, k, a, b)
				}
			}
		}
	}
}

func Curry5[T any, K any, A any, B any, C any, R any](fn func(T, K, A, B, C) R) func(C) func(B) func(A) func(K) func(T) R {
	return func(c C) func(B) func(A) func(K) func(T) R {
		return func(b B) func(A) func(K) func(T) R {
			return func(a A) func(K) func(T) R {
				return func(k K) func(T) R {
					return func(t T) R {
						return fn(t, k, a, b, c)
					}
				}
			}
		}
	}
}
