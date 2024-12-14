package crud.java_crud_api.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import crud.java_crud_api.model.UpdateCompanyModel;
import crud.java_crud_api.repository.UpdateCompanyRepository;


@Service
public class UpdateCompanyService {
     private final UpdateCompanyRepository updateCompanyRepository;

    @Autowired
    public UpdateCompanyService(UpdateCompanyRepository updateCompanyRepository) {
        this.updateCompanyRepository = updateCompanyRepository;
    }

    public UpdateCompanyModel update(UpdateCompanyModel companyUpdate, String emailFromJwt) {
        // Busca a entidade existente no banco de dados
        UpdateCompanyModel existingCompany = updateCompanyRepository.findById(companyUpdate.getId())
                .orElseThrow(() -> new RuntimeException("Usuário não encontrado"));

        if (!existingCompany.getEmail().equals(emailFromJwt)) {
            throw new SecurityException("Usuário não autorizado a atualizar esta conta.");
        }

        // Atualiza os campos do objeto existente
        existingCompany.setNome(companyUpdate.getNome());
        existingCompany.setInscricaoEstadual(companyUpdate.getInscricaoEstadual());
        existingCompany.setDataFundacao(companyUpdate.getDataFundacao());
        existingCompany.setEndereco(companyUpdate.getEndereco());
        existingCompany.setBairro(companyUpdate.getBairro());
        existingCompany.setCidade(companyUpdate.getCidade());
        existingCompany.setEstado(companyUpdate.getEstado());
        existingCompany.setCep(companyUpdate.getCep());
        existingCompany.setTelefone(companyUpdate.getTelefone());
        existingCompany.setSite(companyUpdate.getSite());
        existingCompany.setRegimeTributario(companyUpdate.getRegimeTributario());
        existingCompany.setInscricaoMunicipal(companyUpdate.getInscricaoMunicipal());
        existingCompany.setNaturezaJuridica(companyUpdate.getNaturezaJuridica());
        existingCompany.setNomeFantasia(companyUpdate.getNomeFantasia());
        existingCompany.setRamoAtividade(companyUpdate.getRamoAtividade());
        existingCompany.setCapitalSocial(companyUpdate.getCapitalSocial());
        existingCompany.setQuantidadeFuncionarios(companyUpdate.getQuantidadeFuncionarios());
        existingCompany.setDescricao(companyUpdate.getDescricao());
        existingCompany.setStatus(companyUpdate.getStatus());
        existingCompany.setUltimaAtualizacao(companyUpdate.getUltimaAtualizacao());

        // Salva e retorna a entidade atualizada
        return updateCompanyRepository.save(existingCompany);
    }
    
}
