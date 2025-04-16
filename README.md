# GitHub App Token Generator

このツールは、GitHub Appの認証トークンを生成するためのコマンドラインツールです。

## インストール

```bash
go install github.com/yutachaos/github-app-token@latest
```

## 使用方法

```bash
github-app-token -app-id YOUR_APP_ID -private-key path/to/private-key.pem
```

### オプション

- `-app-id`: GitHub App ID（必須）
- `-private-key`: 秘密鍵ファイルのパス（必須）
- `-expiration`: トークンの有効期限（デフォルト: 10分）

### 例

```bash
# デフォルトの有効期限（10分）でトークンを生成
github-app-token -app-id 123456 -private-key private-key.pem

# カスタム有効期限（1時間）でトークンを生成
github-app-token -app-id 123456 -private-key private-key.pem -expiration 1h
```

## ライセンス

このプロジェクトはMITライセンスの下で公開されています。詳細は[LICENSE](LICENSE)ファイルを参照してください。 