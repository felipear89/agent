# Agent Legislação SP - Frontend

A modern React TypeScript application for managing São Paulo legislation with authentication and best practices.

## 🚀 Features

- **Modern React 19** with TypeScript
- **Authentication System** with context-based state management
- **Form Handling** with react-hook-form
- **Routing** with react-router-dom v7
- **Code Quality** with ESLint and Prettier
- **Build Tool** with Vite
- **Responsive Design** with modern CSS

## 📋 Prerequisites

- Node.js 18+ 
- npm or yarn

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start development server**
   ```bash
   npm run dev
   ```

4. **Open your browser**
   Navigate to `http://localhost:5173`

## 📜 Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build
- `npm run lint` - Run ESLint
- `npm run lint:fix` - Fix ESLint errors
- `npm run type-check` - Run TypeScript type checking
- `npm run format` - Format code with Prettier
- `npm run format:check` - Check code formatting

## 🏗️ Project Structure

```
src/
├── components/          # Reusable UI components
│   └── LoadingSpinner.tsx
├── contexts/           # React contexts
│   └── AuthContext.tsx
├── hooks/              # Custom React hooks
│   └── useAuth.ts
├── lib/                # Utility libraries
│   └── storage.ts
├── pages/              # Page components
│   ├── LoginPage.tsx
│   └── WelcomePage.tsx
├── routes/             # Routing configuration
│   └── index.tsx
├── config/             # Configuration files
│   └── env.ts
├── assets/             # Static assets
├── index.css           # Global styles
└── main.tsx           # Application entry point
```

## 🔧 Configuration

### Environment Variables

Create a `.env` file in the root directory:

```env
VITE_API_URL=http://localhost:3000
```

### TypeScript Configuration

The project uses strict TypeScript configuration with path aliases:

- `@/*` maps to `src/*`
- Strict type checking enabled
- Modern ES2020 target

## 🎨 Styling

The project uses a custom CSS approach with:
- CSS custom properties for theming
- Utility classes for common patterns
- Responsive design principles
- Modern CSS features (Grid, Flexbox, etc.)

## 🔐 Authentication

The authentication system includes:
- Context-based state management
- Local storage persistence
- Loading states
- Error handling
- Protected routes

## 📱 Browser Support

- Chrome 90+
- Firefox 88+
- Safari 14+
- Edge 90+

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support, please open an issue in the repository or contact the development team.
