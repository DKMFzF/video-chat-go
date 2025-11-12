import React, { useState } from "react";
import VideoChat from "../components/VideoChat";
import { generateId } from "../utils/uuid";

const MainPage: React.FC = () => {
  const [joined, setJoined] = useState(false);
  const [userId, setUserId] = useState(generateId());

  return (
    <div>
      {!joined ? (
        <div>
          <h1>0.0.31</h1>
          <button onClick={() => setJoined(true)}>Тыкаем</button>
        </div>
      ) : (
        <VideoChat wsUrl="ws://localhost:8080/ws" userId={userId} />
      )}
    </div>
  );
};

export default MainPage;
