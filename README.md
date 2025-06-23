# Tubely

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Tubely is a modern SaaS platform designed for Creators to efficiently manage, store, and serve their video assets. Built with Go and AWS S3, Tubely ensures robust, scalable, and high-performance file management for creators of all sizes.

---

## 🚀 Features

-   **Efficient Asset Management:** Upload, organize, and serve large video and image files.
-   **AWS S3 Integration:** Secure, scalable cloud storage for all your assets.
-   **Built with Go:** Fast, reliable backend for optimal performance.
-   **Easy Setup:** Minimal configuration to get started.
-   **Open Source:** Contributions welcome!

---

## 📦 Quickstart

### 1. Prerequisites

-   [Go](https://golang.org/doc/install) (v1.21+)
-   [FFMPEG](https://ffmpeg.org/download.html) (`ffmpeg` and `ffprobe` in your `PATH`)
-   [AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)
-   [SQLite 3](https://www.sqlite.org/download.html) (optional, for manual DB inspection)

### 2. Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/your-repo/tubely.git
cd tubely
go mod download
```

Install FFMPEG and SQLite3:

```bash
# Linux
sudo apt update
sudo apt install ffmpeg sqlite3

# macOS
brew update
brew install ffmpeg sqlite3
```

### 3. Download Sample Assets

```bash
./assetdownload.sh
# Creates a samples/ directory with sample images and videos
```

### 4. Configure Environment Variables

Copy the example environment file and update as needed:

```bash
cp .env.example .env
# Edit .env with your AWS and app credentials
```

### 5. Run the Server

```bash
go run .
```

-   A new `tubely.db` database and `assets` directory will be created.
-   Visit the link in your console to access the local web page.

---

## 🛠️ Usage Example

Upload a video asset via the API:

```bash
curl -X POST -F "file=@/path/to/video.mp4" http://localhost:8080/upload
```

List all uploaded assets:

```bash
curl http://localhost:8080/assets
```

---

## 🤝 Contributing

We welcome contributions! To get started:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature`
3. Make your changes and commit: `git commit -m 'Add new feature'`
4. Push to your fork: `git push origin feature/your-feature`
5. Open a Pull Request.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## 🙋‍♂️ Support & Community

-   [Boot.dev Course](https://boot.dev/courses/learn-file-storage-golang)
-   [Open an Issue](https://github.com/your-repo/tubely/issues) for bugs or feature requests.

---

> Tubely: Empowering YouTubers with scalable, open source asset management solutions.
