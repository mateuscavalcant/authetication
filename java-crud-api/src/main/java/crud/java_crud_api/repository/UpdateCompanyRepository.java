package crud.java_crud_api.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import crud.java_crud_api.model.UpdateCompanyModel;

public interface UpdateCompanyRepository extends JpaRepository<UpdateCompanyModel, Long> {
    UpdateCompanyModel findById(long id);
}
