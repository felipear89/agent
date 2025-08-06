import { useAuthContext } from '@/contexts/AuthContext';
import './DemoPage.css';

export default function DemoPage() {
  const { logout } = useAuthContext();
  const script = document.querySelector<HTMLIFrameElement>('iframe[src*="agents"]');
  if (script) {
    script.style.display = '';
  }

  return (
    <div className="page-container">
      <div className="demo-container">
        <script
          async
          src="https://ot7q6l7i6k57sxp2ikoi72vx.agents.do-ai.run/static/chatbot/widget.js"
          data-agent-id="46aa25db-72bf-11f0-b074-4e013e2ddde4"
          data-chatbot-id="0k-lYmsjp4-blLxB6U2qzWkfaZk7drx1"
          data-name="agent-legislacao Chatbot"
          data-primary-color="#031B4E"
          data-secondary-color="#E5E8ED"
          data-button-background-color="#0061EB"
          data-starting-message="Olá eu sou a SMULLover, como posso te ajudar ?"
          data-logo="/static/chatbot/icons/default-agent.svg"
        ></script>
        <button className="logout-button-demo" onClick={logout}>
          Log Out
        </button>
      </div>
    </div>
  );
}
