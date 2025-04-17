# ResumaAI

**ResumaAI** is an open-source, AI-enhanced resume builder designed for speed, simplicity, and extensibility. Built with a minimalistic interface and modern tech stack, it helps users craft clean, professional resumes with smart suggestions powered by OpenAI.

## ✨ Features

- Minimalistic UI with dark theme (Next.js + Tailwind CSS)
- AI-powered content suggestions using OpenAI
- Resume editing with structured fields
- Export to PDF and Word formats
- Multiple professional templates
- Optional user authentication
- Modular, scalable backend in Go

## 🧠 Tech Stack

- **Frontend:** Next.js 13+ (App Router), Tailwind CSS, TypeScript
- **Backend:** Go (REST API), Gorilla Mux, OpenAI API integration
- **AI:** OpenAI GPT (content enhancement, keyword optimization)
- **Export:** PDF generation via Go (or frontend export options)
- **Deployment:** Vercel (frontend), Fly.io / Railway / Render (backend)

## 🚀 Getting Started

### Frontend

```bash
cd web
npm install
npm run dev
```

Access the frontend at [http://localhost:3000](http://localhost:3000)

### Backend

```bash
cd server
go run ./cmd/api
```

Ensure you have a `.env` file in `server/` with:

```
OPENAI_API_KEY=your_openai_key_here
```

API available at [http://localhost:8000](http://localhost:8000)

## 📦 Project Structure

```
resume-builder/
├── web/      # Next.js frontend
├── server/   # Go backend
```

## 🛠 For Contributors

We welcome contributions! Ideal areas to contribute:
- Add new resume templates
- Improve AI prompt handling
- Refactor components or add new features
- Help with internationalization or accessibility

## 📄 License

MIT License. See [LICENSE](./LICENSE) for more info.
