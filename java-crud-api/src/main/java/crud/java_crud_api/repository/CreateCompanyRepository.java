package crud.java_crud_api.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import crud.java_crud_api.model.CreateCompanyModel;

public interface CreateCompanyRepository extends JpaRepository<CreateCompanyModel, Long> {
    boolean existByEmail(String email);
    boolean existByCNPJ(String cnpj);
}
