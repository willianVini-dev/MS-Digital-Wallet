# 💳 Carteira Digital - Microserviços (Projeto de Aprendizado)

Este é um **projeto de aprendizado** que simula uma Carteira Digital utilizando arquitetura de **microserviços** em **Go**, com frontend em **React**.  
O objetivo é praticar conceitos de sistemas distribuídos, comunicação síncrona e assíncrona, banco por serviço e orquestração de pagamentos.

---

## 🧩 Arquitetura

O projeto segue uma arquitetura inspirada em C4/PlantUML:

- **API (React)**: interface web do usuário, faz chamadas HTTP/JSON para os microsserviços.
- **Wallet Core (Go)**: serviço central que orquestra transações e publica eventos assíncronos.
- **Balance Service (Go)**: gerencia saldo da carteira e atualizações financeiras.
- **Statement Service (Go)**: gera e consulta extratos, consome eventos do Wallet Core.
- **Payment ACL (Go)**: camada de antifraude e integração com gateway de pagamento externo.
- **Bancos de dados (MySQL)**: cada serviço possui seu próprio banco, garantindo isolamento.

---

## 🔄 Comunicação

- **Síncrona**: API → microsserviços (HTTP/JSON), Wallet Core → Payment ACL → Payment Gateway  
- **Assíncrona**: Wallet Core → Balance Service / Statement Service (event-driven)

---
