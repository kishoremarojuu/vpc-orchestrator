# VPC Orchestrator

A Go-based CLI tool to automate AWS VPC peering across multiple accounts. Designed with modular orchestration layers to handle creation, polling, and validation of VPC peering connections.

---

## 🔧 Features
- Create VPC peering connections via CLI
- Poll for peering readiness
- Modular design for future support (route table, security group updates)
- Built with Go using AWS SDK v2 and Cobra

---

## 🚀 Usage

### CLI Example
```bash
vpc-orchestrator peer \
  --source-vpc vpc-12345678 \
  --target-vpc vpc-98765432 \
  --peer-account 123456789012
```

### Flags
- `--source-vpc`: Source VPC ID
- `--target-vpc`: Target VPC ID
- `--peer-account`: AWS Account ID of the target VPC

---

## 📁 Project Structure
```
vpc-orchestrator/
├── main.go                    # Entry point
├── cmd/                      # CLI command handlers
│   └── peer.go
├── orchestrator/            # Peering orchestration logic
│   └── peering.go
├── utils/                   # AWS client helpers
│   └── aws.go
├── go.mod / go.sum          # Go modules
└── README.md
```

---

## 🧠 How It Works
1. User triggers the CLI with VPC details.
2. Go program uses AWS SDK v2 to:
    - Create the peering connection
    - Poll until it becomes active
    - (Optional: Update route tables & security groups — coming soon)
3. All operations logged to stdout.

---

## 📦 Dependencies
- Go 1.18+
- [AWS SDK v2 for Go](https://github.com/aws/aws-sdk-go-v2)
- [Cobra CLI](https://github.com/spf13/cobra)

---

## 🔒 Requirements
- AWS credentials must be configured (e.g., `~/.aws/credentials` or environment variables)
- Proper IAM permissions to manage VPCs, routes, and security groups

---

## 📌 Roadmap
- [ ] Route Table updates
- [ ] Security Group updates
- [ ] Logging and metrics with Zap
- [ ] Docker support

---

