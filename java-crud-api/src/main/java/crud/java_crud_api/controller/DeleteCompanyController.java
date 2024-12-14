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
import crud.java_crud_api.model.DeleteCompanyModel;
import crud.java_crud_api.service.DeleteCompanyService;


@RestController
@RequestMapping("/api")
@CrossOrigin("")
public class DeleteCompanyController {

    private final DeleteCompanyService companyService;
    private final RegisterJwt registerJwt;

    @Autowired
    public DeleteCompanyController(DeleteCompanyService companyService, RegisterJwt registerJwt) {
        this.companyService = companyService;
        this.registerJwt = registerJwt;
    }

    @PostMapping("/company/delete")
    public ResponseEntity<Map<String, String>> delete(
            @RequestBody DeleteCompanyModel deleteRegister,
            @RequestHeader(value = "Authorization", required = false) String authorizationHeader) {

        Map<String, String> response = new HashMap<>();

        // Checar Authorization header
        if (authorizationHeader == null || !authorizationHeader.startsWith("Bearer ")) {
            response.put("error", "Authorization header is missing or invalid.");
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(response);
        }

        String jwtToken = authorizationHeader.substring(7);
        String emailFromJwt;

        // Extrair email do JWT
        try {
            emailFromJwt = registerJwt.extractEmail(jwtToken);
        } catch (Exception e) {
            response.put("error", "Invalid JWT token.");
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body(response);
        }

        try {
            // Chamar delete service
            companyService.delete(deleteRegister, emailFromJwt, deleteRegister.getEmail(), deleteRegister.getPassword());
            response.put("message", "Cadastro deletado com sucessso");
            return ResponseEntity.status(HttpStatus.OK).body(response);
        } catch (SecurityException e) {
            response.put("error", "Unauthorized access: " + e.getMessage());
            return ResponseEntity.status(HttpStatus.FORBIDDEN).body(response);
        } catch (RuntimeException e) {
            response.put("error", "Operation failed: " + e.getMessage());
            return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(response);
        }
    }
}