package crud.java_crud_api.model;

import java.time.LocalDate;

public class CreateCompanyModel extends CompanyModel {

    public CreateCompanyModel() {}

    public CreateCompanyModel(Long id,
    String nome,
    String password,
    String cnpj,
    LocalDate dataFundacao,
    String endereco,
    String bairro,
    String cidade,
    String estado,
    String cep,
    String telefone,
    String email
) {
    this.id = id;
    this.nome = nome;
    this.password = password;
    this.cnpj = cnpj;
    this.dataFundacao = dataFundacao;
    this.endereco = endereco;
    this.bairro = bairro;
    this.cidade = cidade;
    this.estado = estado;
    this.cep = cep;
    this.telefone = telefone;
    this.email = email;
}
    
}
