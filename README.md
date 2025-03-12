# Security Bank Password

Este é um projeto de exemplo que utiliza o framework Gin para criar uma API de validação de senhas baseada em pares de números e um dígito validador.

Aqui está um tutorial passo a passo para instalar o Go no seu computador.  

---

# 📌 Como instalar Go no seu computador  

## 🖥️ Passo 1: Baixar o Go  
1. Acesse o site oficial do Go: [https://go.dev/dl/](https://go.dev/dl/)  
2. Baixe a versão mais recente do Go compatível com seu sistema operacional:  
   - **Windows:** Arquivo `.msi`  
   - **Linux:** Arquivo `.tar.gz`  
   - **macOS:** Arquivo `.pkg`  

## 🏗️ Passo 2: Instalar o Go  

### 🔹 No Windows  
1. Execute o arquivo `.msi` baixado.  
2. Siga as instruções do instalador.  
3. Ao finalizar, abra o **Prompt de Comando (cmd)** ou **PowerShell** e execute:  
   ```sh
   go version
   ```
   Se a instalação foi bem-sucedida, você verá a versão instalada do Go.  

### 🔹 No Linux  
1. Abra o terminal e navegue até a pasta onde baixou o arquivo.  
2. Extraia o arquivo para `/usr/local` com o comando:  
   ```sh
   sudo tar -C /usr/local -xzf go1.xx.linux-amd64.tar.gz
   ```
3. Adicione o Go ao seu PATH no arquivo `~/.bashrc` ou `~/.profile`:  
   ```sh
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```
4. Teste a instalação executando:  
   ```sh
   go version
   ```

### 🔹 No macOS  
1. Execute o arquivo `.pkg` baixado.  
2. Siga as instruções do instalador.  
3. Verifique se a instalação foi concluída corretamente com:  
   ```sh
   go version
   ```

## 📁 Passo 3: Configurar o GOPATH (Opcional)  
O **GOPATH** define onde os pacotes e projetos Go serão armazenados.  

1. Crie uma pasta para armazenar os projetos Go (por exemplo, `~/go`):  
   ```sh
   mkdir -p ~/go
   ```
2. Adicione essa pasta ao ambiente do sistema:  
   - **Linux/macOS:**  
     ```sh
     echo 'export GOPATH=$HOME/go' >> ~/.bashrc
     echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
     source ~/.bashrc
     ```
   - **Windows:**  
     - Abra as **Configurações do Sistema** → **Variáveis de Ambiente**  
     - Adicione uma nova variável chamada `GOPATH` com o valor `C:\Users\SeuUsuário\go`  
     - Edite a variável `Path` e adicione `%GOPATH%\bin`  

Agora você tem o Go instalado e configurado no seu computador!

## Funcionalidades

- Geração de pares de números.
- Geração de dígito validador.
- Validação de senhas enviadas pelo usuário.

## Endpoints

### GET /pairs

Gera e retorna pares de números e um dígito validador.

**Resposta:**

```json
{
  "pairs": [[1, 2], [3, 4], [5, 6], [7, 8], [9, 0]],
  "validator": 42
}
```
### POST /login

Valida a senha enviada pelo usuário.

**Requisição:**
**Requisição:**

```json
{
  "pairs": [[1, 2], [3, 4], [5, 6], [7, 8], [9, 0]],
  "validator": 42,
  "password": [[1, 2], [3, 4], [5, 6], [7, 8], [9, 0]]
}
```

**Resposta:**

**Sucesso:**

```json
{
  "message": "Welcome!"
}
```

**Erro:**

```json
{
  "error": "Unauthorized"
}
```

## Como Executar

1. Clone o repositório:
   ```sh
   git clone <URL_DO_REPOSITORIO>
   ```
2. Navegue até o diretório do projeto:
   ```sh
   cd security-bank-password
   ```
3. Instale as dependências:
   ```sh
   go mod download
   ```
4. Execute o projeto:
   ```sh
   go run main.go
   ```

## Estrutura do Projeto

- `main.go`: Arquivo principal que define os endpoints e a lógica de validação.
- `password/pair-generator.go`: Contém funções para gerar pares de números, dígito validador e validar senhas.

## Dependências

- **Gin**: Framework web para Go.
