package crud.java_crud_api.controller;

import java.util.HashMap;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import crud.java_crud_api.jwt.RegisterJwt;
import crud.java_crud_api.model.UpdateCompanyModel;
import crud.java_crud_api.service.UpdateCompanyService;

@RestController
@RequestMapping("/api")
@CrossOrigin("")
public class UpdateCompanyController {

    private final UpdateCompanyService companyService;
    private final RegisterJwt registerJwt;

    @Autowired
    public UpdateCompanyController(UpdateCompanyService companyService, RegisterJwt registerJwt) {
        this.companyService = companyService;
        this.registerJwt = registerJwt;
    }

    @PostMapping("/company/update")
    public ResponseEntity<Map<String, String>> update(
            @RequestBody UpdateCompanyModel updateRegister,
            @RequestHeader(value = "Authorization", required = false) String authorizationHeader) {
        Map<String, String> response = new HashMap<>();

        if (authorizationHeader == null || !authorizationHeader.startsWith("Bearer ")) {
            response.put("error", "Authorization header is missing or invalid.");
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(response);
        }

        String jwtToken = authorizationHeader.substring(7);
        String email;
        try {
            email = registerJwt.extractEmail(jwtToken);
        } catch (Exception e) {
            response.put("error", "Invalid JWT token.");
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(response);
        }

        try {
            companyService.update(updateRegister, email);
            response.put("message", "Cadastro atualizado com sucesso.");
            return ResponseEntity.status(HttpStatus.OK).body(response);
        } catch (Exception e) {
            response.put("error", e.getMessage());
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(response);
        }
    }
}
