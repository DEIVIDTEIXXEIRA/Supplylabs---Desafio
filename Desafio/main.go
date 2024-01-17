package main

// RuleSet representa um conjunto de regras entre opções.
type RuleSet struct {
	dependencias map[string]string
	conflitos    map[string]string
}

// NewRuleSet retorna um conjunto de regras vazio.
func NewRuleSet() *RuleSet {
	return &RuleSet{
		dependencias: make(map[string]string),
		conflitos:    make(map[string]string),
	}
}

// AddDep adiciona uma nova dependência entre A e B.
func (rs *RuleSet) AddDep(A, B string) {
	rs.dependencias[A] = B
}

// AddConflict adiciona um novo conflito entre A e B.
func (rs *RuleSet) AddConflict(A, B string) {
	rs.conflitos[A] = B
	rs.conflitos[B] = A
}

// IsCoherente verifica se o conjunto de regras é coerente.
func (rs *RuleSet) IsCoherent() bool {
	visitado := make(map[string]bool)
	naPilha := make(map[string]bool)

	// função recursiva que verifica se a opção esta na pilha indicando um ciclo, se ja foi visitada, não há ciclos.
	var temCiclo func(string) bool
	temCiclo = func(opcao string) bool {
		if naPilha[opcao] {
			return true
		}
		if visitado[opcao] {
			return false
		}

		// atribui a opção ja vista naPilha
		visitado[opcao] = true
		naPilha[opcao] = true

		// obtem as dependencias e conflitos da opção atual
		dependencia, existeDependencia := rs.dependencias[opcao]
		conflito, existeConflito := rs.conflitos[opcao]

		// verfica se ha dependencia e se essa dependencia tem ciclo.
		if existeDependencia && temCiclo(dependencia) {
			return true
		}

		// verifica se ha conflito e se houver, verifica se esta na pilha e se há um ciclo no conflito
		if existeConflito {
			if naPilha[conflito] {
				return true
			}
			if temCiclo(conflito) {
				return true
			}
		}

		naPilha[opcao] = false

		return false
	}

	for opcao := range rs.dependencias {
		if temCiclo(opcao) {
			return false
		}
	}

	return true
}
