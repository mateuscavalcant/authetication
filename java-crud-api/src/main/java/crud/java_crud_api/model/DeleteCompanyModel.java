package crud.java_crud_api.model;

public class DeleteCompanyModel extends CompanyModel {

    public DeleteCompanyModel(){}

    public DeleteCompanyModel(Long id, String email, String password) {
        this.id = id;
        this.email = email;
        this.password = password;
    }
    
}
