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
        <script async
  src="https://hq4istavj2q6qaw22ivdp7pi.agents.do-ai.run/static/chatbot/widget.js"
  data-agent-id="f25433a6-0081-11f1-b074-4e013e2ddde4"
  data-chatbot-id="-ETDbqzZ2sQE-l3MJSI5I9Ph1vE5Fj56"
  data-name="agent-SANEAGO Chatbot"
  data-primary-color="#031B4E"
  data-secondary-color="#E5E8ED"
  data-button-background-color="#0061EB"
  data-starting-message="Eu sou seu assistente para documentação de PPP de Esgotos da SANEAGO, em posso te ajudar"
  data-logo="/static/chatbot/icons/default-agent.svg">
</script>
        <button className="logout-button-demo" onClick={logout}>
          Log Out
        </button>
      </div>
    </div>
  );
}
