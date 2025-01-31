![image](https://github.com/user-attachments/assets/e6a71825-9367-4f6c-b10c-bc2555e934cc)

# Poc Go EKS

## Descrição

projeto de prova de conceito que demonstra uma aplicação em Go implantada no Amazon EKS (Elastic Kubernetes Service). A aplicação fornece endpoints HTTP para escrever e ler mensagens de um banco de dados MySQL, utilizando Kubernetes para orquestração e Docker para containerização.

## Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gorilla Mux](https://github.com/gorilla/mux)
- [MySQL](https://www.mysql.com/)
- [Docker](https://www.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Amazon EKS](https://aws.amazon.com/eks/)
- [eksctl](https://eksctl.io/)

## Estrutura do Projeto

### Estrutura de Diretórios

```
/cmd
    main.go
/internal
  /handlers
    read.go
    write.go
/k8s
  mysql-master.yaml
  mysql-slave.yaml
  poc-go-eks.yaml
  secrets.yaml
README.md
Dockerfile
```

### Requisitos

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started)
- [Kubernetes CLI (kubectl)](https://kubernetes.io/docs/tasks/tools/)
- [AWS CLI](https://aws.amazon.com/cli/)
- [eksctl](https://eksctl.io/)

### Instalação

1. Clone o repositório:

  ```bash
  git clone https://github.com/seu-usuario/poc-go-eks.git
  cd poc-go-eks
  ```

2. Construir a imagem Docker:

  ```bash
  docker build -t poc-go-eks .
  ```

3. Realize o push da imagem no ECR:

  ```bash
  docker push <ID_DA_CONTA_AWS>.dkr.ecr.<região>.amazonaws.com/poc-go-eks:latest
  ```

4. Configurar o cluster com eksctl:

  ```bash
  eksctl create cluster --name poc-go-cluster --region us-west-2 --nodegroup-name poc-nodes --node-type t2.medium --nodes 3
  ```

5. Implantar no Amazon EKS:

  ```bash
    kubectl apply -f k8s/mysql-master.yaml
    kubectl apply -f k8s/mysql-slave.yaml
    kubectl apply -f k8s/poc-go-eks.yaml
    kubectl apply -f k8s/secrets.yaml
  ```

### Uso

Após a implantação, a aplicação estará disponível através dos endpoints HTTP configurados. Por exemplo:

- **Escrever mensagem**: `POST /messages`

  Corpo da requisição:

  ```json
  {
    "message": "Sua mensagem aqui"
  }
  ```

- **Ler mensagens**: `GET /messages`
