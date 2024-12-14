package crud.java_crud_api.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import crud.java_crud_api.model.DeleteCompanyModel;

public interface DeleteCompanyRepository extends JpaRepository<DeleteCompanyModel, Long> {
    DeleteCompanyModel findById(long id);
    DeleteCompanyModel findByEmail(String email);
}
