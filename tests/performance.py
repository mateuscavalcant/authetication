import unittest
import requests

class TestSignupEndpoint(unittest.TestCase):

    def test_signup_success(self):
        url = 'http://localhost:8000/signup'
        for i in range(1, 11):
            # Gerar um endereço de email único
            email = f"xteste{i}@example.com"

            data = {
                'name': 'User Test',
                'email': email,
                'password': 'password123',
                'confirm_password': 'password123'
            }
            response = requests.post(url, data=data)
            self.assertEqual(response.status_code, 200)


if __name__ == '__main__':
    unittest.main()
