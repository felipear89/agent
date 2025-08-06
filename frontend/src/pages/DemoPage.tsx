import { useAuthContext } from '@/contexts/AuthContext';
import './DemoPage.css';

export default function DemoPage() {
  const { logout } = useAuthContext();

  return (
    <div className="page-container">
      <div className="demo-container">
        <script
          async
          src="https://wdtibedlmha6tbwcohjooecj.agents.do-ai.run/static/chatbot/widget.js"
          data-agent-id="999166d5-723c-11f0-b074-4e013e2ddde4"
          data-chatbot-id="gPVgoAIfKHHwBeM46Qzo1pLa0Mr9D2Rm"
          data-name="Kubernetes Genius Chatbot"
          data-primary-color="#031B4E"
          data-secondary-color="#E5E8ED"
          data-button-background-color="#0061EB"
          data-starting-message="Hello! How can I help you today?"
          data-logo="/static/chatbot/icons/default-agent.svg"
        ></script>
        <button className="logout-button-demo" onClick={logout}>
          Log Out
        </button>
      </div>
    </div>
  );
}
