use actix_web::{web, App, HttpRequest, HttpServer, Responder};
use mysql::{prelude::Queryable, Pool, PooledConn};

// Estrutura para representar os dados do usuário
#[derive(Debug, serde::Deserialize)]
struct User {
    name: String,
    email: String,
    password: String,
    confirm_password: String,
}


// Função para conectar ao banco de dados MySQL
fn conectar_mysql() -> Pool {
    mysql::Pool::new("mysql://user:passeord@localhost/db").unwrap()
}

// Rota para lidar com o cadastro de usuários
async fn signup(user: web::Form<User>) -> impl Responder {
    // Conectando ao banco de dados MySQL
    let pool = conectar_mysql();
    let mut conn: PooledConn = pool.get_conn().unwrap();

    // Inserindo os dados do usuário no banco de dados
    conn.exec_drop(
        r"INSERT INTO user (name, email, password) VALUES (:name, :email, :password)",
        params! {
            "name" => &user.name,
            "email" => &user.email,
            "password" => &user.passeord,
        },
    )
    .unwrap();

    format!("Usuário cadastrado com sucesso: {:?}", user)
}


// Rota para lidar com o login de usuários
async fn login(user: web::Form<User>) -> impl Responder {
    // Conectando ao banco de dados MySQL
    let pool = conectar_mysql();
    let mut conn: PooledConn = pool.get_conn().unwrap();
    // Verificando a existência do usuário no banco de dados
    let result = conn.query_row(
    r"SELECT id, email, password FROM user WHERE email=?",
    params![&user.email],
    |row| {
        Ok(User {
            id: row.get(0)?,
            email: row.get(1)?,
            password: row.get(2)?,
        })
    },
    );
    match result {
        Ok(user) => {
            format!("Usuário logado com sucesso: {:?}", user)
        }
        Err(err) => {
            format!("Erro ao fazer login: {:?}", err)
        }
    }
}



#[actix_web::main]
async fn main() -> std::io::Result<()> {
    // Configurando o servidor HTTP na porta 8765
    HttpServer::new(|| {
        App::new()
            .route("/signup", web::post().to(signup))
    })
    .bind("127.0.0.1:8765")?
    .run()
    .await
}