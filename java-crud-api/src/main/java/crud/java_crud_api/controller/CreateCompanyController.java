package crud.java_crud_api.controller;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import crud.java_crud_api.model.CreateCompanyModel;
import crud.java_crud_api.service.CreateCompanyService;


@RestController
@RequestMapping("/api/auth")
@CrossOrigin("")
public class CreateCompanyController {
     private final CreateCompanyService companyService;


    @Autowired
    public CreateCompanyController(CreateCompanyService companyService) {
        this.companyService = companyService;
    }


    @PostMapping("/enterprise/signup")
    public ResponseEntity<Map<String, String>> signup(@RequestBody CreateCompanyModel createRegister) {
        Map<String, String> response = new HashMap<>();

        try {
        companyService.signup(createRegister);
        response.put("message", "Cadastro realizado com sucesso.");
        return ResponseEntity.status(HttpStatus.CREATED).body(response);

        } catch (Exception e) {
            response.put("error", e.getMessage());
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(response);
        }

    }
    
    
}
