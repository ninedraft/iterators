package iterators

func none[E any]() (E, bool) {
	return empty[E](), false
}

func empty[E any]() E {
	var empty E
	return empty
}
