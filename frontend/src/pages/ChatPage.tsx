import { useRef, useState } from 'react';
import { useAuthContext } from '@/contexts/AuthContext';
import './ChatPage.css';

interface Message {
  id: number;
  text: string;
  sender: 'user' | 'other';
  timestamp: Date;
}

interface ChatUser {
  id: number;
  name: string;
  avatar: string;
  isOnline: boolean;
}

export default function ChatPage() {
  const { logout } = useAuthContext();
  const [messages, setMessages] = useState<Message[]>([
    {
      id: 1,
      text: "Olá! Como você está?",
      sender: 'other',
      timestamp: new Date(Date.now() - 300000) // 5 minutes ago
    },
    {
      id: 2,
      text: "Oi! Estou bem, obrigado! E você?",
      sender: 'user',
      timestamp: new Date(Date.now() - 240000) // 4 minutes ago
    },
    {
      id: 3,
      text: "Também estou bem! Que tal aquele projeto que estávamos discutindo?",
      sender: 'other',
      timestamp: new Date(Date.now() - 180000) // 3 minutes ago
    },
    {
      id: 4,
      text: "Está indo muito bem! Já implementei a funcionalidade principal.",
      sender: 'user',
      timestamp: new Date(Date.now() - 120000) // 2 minutes ago
    },
    {
      id: 5,
      text: "Que ótimo! Posso ver o código quando estiver pronto?",
      sender: 'other',
      timestamp: new Date(Date.now() - 60000) // 1 minute ago
    }
  ]);

  const [newMessage, setNewMessage] = useState('');
  const nextMessageId = useRef(6);

  const mockUsers: ChatUser[] = [
    {
      id: 1,
      name: "João Silva",
      avatar: "👨‍💻",
      isOnline: true
    },
    {
      id: 2,
      name: "Maria Santos",
      avatar: "👩‍💼",
      isOnline: true
    },
    {
      id: 3,
      name: "Pedro Costa",
      avatar: "👨‍🔧",
      isOnline: false
    },
    {
      id: 4,
      name: "Ana Oliveira",
      avatar: "👩‍🎨",
      isOnline: true
    }
  ];

  const handleSendMessage = () => {
    if (newMessage.trim()) {
      const message: Message = {
        id: nextMessageId.current,
        text: newMessage,
        sender: 'user',
        timestamp: new Date()
      };
      
      setMessages(prev => [...prev, message]);
      setNewMessage('');
      nextMessageId.current += 1;

      const m: Message = {
        id: nextMessageId.current,
        text: '...',
        sender: 'other',
        timestamp: new Date()
      }
      setMessages(prev => [...prev, m]);

      // Simulate response after 1 second
      setTimeout(() => {
        const responses = [
          "Interessante! Me conte mais sobre isso.",
          "Entendi perfeitamente!",
          "Que legal! Estou ansioso para ver o resultado.",
          "Ótima ideia! Como posso ajudar?",
          "Perfeito! Vamos continuar assim."
        ];
        
        const randomResponse = responses[Math.floor(Math.random() * responses.length)];
        const responseMessage: Message = {
          id: nextMessageId.current + 1,
          text: randomResponse,
          sender: 'other',
          timestamp: new Date()
        };
        
        setMessages(prev => [...prev.slice(0, -1), responseMessage]);
        //nextMessageId.current+=2;
      }, 2000);
    }
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  const formatTime = (date: Date) => {
    return date.toLocaleTimeString('pt-BR', { 
      hour: '2-digit', 
      minute: '2-digit' 
    });
  };

  return (
    <div className="page-container">
      <div className="chat-container">
        {/* Header */}
        <div className="chat-header">
          <h1>Chat</h1>
          <button className="logout-button" onClick={logout}>
            Log Out
          </button>
        </div>

        <div className="chat-content">

          {/* Chat Messages */}
          <div className="chat-main">
            <div className="messages-container">
              {messages.map(message => (
                <div 
                  key={message.id} 
                  className={`message ${message.sender === 'user' ? 'user-message' : 'other-message'}`}
                >
                  <div className="message-content">
                    <p className="message-text">{message.text}</p>
                    <span className="message-time">
                      {formatTime(message.timestamp)}
                    </span>
                  </div>
                </div>
              ))}
            </div>

            {/* Message Input */}
            <div className="message-input-container">
              <textarea
                className="message-input"
                value={newMessage}
                onChange={(e) => setNewMessage(e.target.value)}
                onKeyDown={handleKeyPress}
                placeholder="Digite sua mensagem..."
                rows={1}
              />
              <button 
                className="send-button"
                onClick={handleSendMessage}
                disabled={!newMessage.trim()}
              >
                Enviar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
} 