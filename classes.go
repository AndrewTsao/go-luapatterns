package luapatterns

func islower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isupper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isalpha(b byte) bool {
	return islower(b) || isupper(b)
}

func iscntrl(b byte) bool {
	return b < ' ' || (b >= '\x7F' && b <= '\x9f')
}

func isdigit(b byte) bool {
	return b >= 48 && b <= 57
}

func ispunct(b byte) bool {
	return	(b >= '{' && b <= '~') || (b == '`') || (b >= '[' && b <= '_') || (b == '@') || (b >= ':' && b <= '?') || (b >= '(' && b <= '/') || (b >= '!' && b <= '\'')
}

func isspace(b byte) bool {
	return b == '\t' || b == '\n' || b == '\v' || b == '\f' || b == '\r' || b == ' '
}

func isalnum(b byte) bool {
	return isalpha(b) || isdigit(b)
}

func isxdigit(b byte) bool {
	return isdigit(b) || (b >= 'a' && b <= 'f') || (b >= 'A' && b <= 'F')
}
