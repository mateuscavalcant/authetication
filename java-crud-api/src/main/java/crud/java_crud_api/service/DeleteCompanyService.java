package crud.java_crud_api.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Service;

import crud.java_crud_api.model.DeleteCompanyModel;
import crud.java_crud_api.repository.DeleteCompanyRepository;

@Service
public class DeleteCompanyService {
    private final DeleteCompanyRepository deleteCompanyRepository;
    private BCryptPasswordEncoder passwordEncoder;

    @Autowired
    public DeleteCompanyService(DeleteCompanyRepository deleteCompanyRepository, BCryptPasswordEncoder passwordEncoder) {
        this.deleteCompanyRepository = deleteCompanyRepository;
        this.passwordEncoder = passwordEncoder;
    }

    public void delete(DeleteCompanyModel companyDelete, String emailFromJwt, String email, String password) {
        // Busca a entidade existente no banco de dados
        DeleteCompanyModel existingCompany = deleteCompanyRepository.findById(companyDelete.getId())
                .orElseThrow(() -> new RuntimeException("Usuário não encontrado"));

        if (!existingCompany.getEmail().equals(emailFromJwt)) {
            throw new SecurityException("Usuário não autorizado a atualizar esta conta.");
        }

        if (deleteCompanyRepository.findByEmail(email) == null) {
            System.out.println("Email não encontrado");
            throw new RuntimeException("Invalid credentials");

        }

        if (!passwordEncoder.matches(password, existingCompany.getPassword())) {
            System.out.println("Failed in encode password");
            throw new RuntimeException("Invalid credentials");
        }

        deleteCompanyRepository.delete(existingCompany);
    }
    
}
