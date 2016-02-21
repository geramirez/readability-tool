package scorer

import "testing"

type test struct {
	sentence string
	expected int
}

type testRunes struct {
	letter   rune
	expected int
}

var countSentencesTests = []test{
	{"Es una receta que puede prepararse en menos de 30 minutos.", 1},
	{"Ponle un tomate grande. Ponle una taza de pasta de tomate", 2},
	{"¡Ideal para mimar a toda la familia!", 1}, {"¿Como se llama?", 1},
	{"¡Hola! ¿Como? Si.", 3}, {"¡Hola!!!! ¿Como??? Si.....", 3},
}

func TestCountSentences(t *testing.T) {
	for _, example := range countSentencesTests {
		actual := countSentences(example.sentence)
		if actual != example.expected {
			t.Errorf("countSentences(%s): expected %d, actual %d", example.sentence, example.expected, actual)
		}
	}
}

var countWordTests = []test{
	{"Es una receta que puede prepararse en menos de 30 minutos.", 11},
	{"Ponle un tomate grande. Ponle una taza de pasta de tomate", 11},
	{"¡Ideal para mimar a toda la familia!", 7}, {"¿Como se llama?", 3},
	{"¡Hola! ¿Como? Si.", 3}, {"¡Hola!!!! ¿Como??? Si.....", 3},
}

func TestGetWords(t *testing.T) {
	for _, example := range countWordTests {
		actual := len(getWords(example.sentence))
		if actual != example.expected {
			t.Errorf("len(getWords(%s)): expected %d, actual %d", example.sentence, example.expected, actual)
		}
	}
}

var getLetterTypeTest = []testRunes{
	{'i', 1}, {'u', 1},
	{'a', 2}, {'e', 2}, {'o', 2}, {'í', 2}, {'ú', 2},
	{'b', 3}, {'c', 3}, {'q', 3}, {'r', 3}, {'z', 3},
}

func TestgetLetterType(t *testing.T) {
	for _, example := range getLetterTypeTest {
		actual := getLetterType(example.letter)
		if actual != example.expected {
			t.Errorf("Got %s, Expected %s", actual, example.expected)
		}
	}
}

var countSyllablesTest = []test{
	{"ola", 2}, {"Amor", 2}, {"Usual", 2}, {"caos", 2}, {"leo", 2},
	{"Traer", 2}, {"aire", 2}, {"ciudad", 2}, {"pie", 2}, {"tomate", 3},
	{"amigo", 3}, {"camiseta", 4}, {"acción", 2}, {"también", 2}, {"cansado", 3},
	{"frío", 2}, {"hablar", 2}, {"problema", 3}, {"chico", 2},
	{"tortilla", 3}, {"Arroz", 2}, {"increíble", 4}, {"improvisar", 4},
	{"Esdrújula", 4}, {"constante", 3}, {"inspector", 3}, {"experto", 3},
	{"extraordinario", 6}, {"inscribir", 3}, {"delicadeza", 5}, {"sinceridad", 4},
	{"teología", 5}, {"último", 3}, {"representando", 5}, {"muerte", 2}, {"también", 2},
	{"cuentista", 3}, {"título", 3}, {"hispanoamericano", 8}, {"noticia", 3},
	{"manumisión", 4}, {"desgraciadamente", 6}, {"nacionalidades", 6}, {"bautizaría", 5},
	{"obstrucción", 3}, {"sintió", 2}, {"irma", 2}, {"mia", 2},
}

func TestCountSyllables(t *testing.T) {
	for _, example := range countSyllablesTest {
		actual := countSyllables(example.sentence)
		if actual != example.expected {
			t.Errorf("countSyllables(%s): expected %d, actual %d", example.sentence, example.expected, actual)
		}

	}
}
