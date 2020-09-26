package databases

// Factory função que cria uma conexao com o banco de dados de acordo com a opcao escolhida
func Factory(t int) DBConnection {
	switch t {
	case 1:
		return &Mysql{}
	case 2:
		return &Postgres{}
	default:
		return nil
	}
}
