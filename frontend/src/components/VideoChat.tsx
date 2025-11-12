import React, { useEffect, useRef, useState } from "react";

type PeerConnections = {
  [key: string]: RTCPeerConnection;
};

type VideoElements = {
  [key: string]: HTMLVideoElement;
};

interface VideoChatProps {
  wsUrl: string;
  userId: string;
}

const VideoChat: React.FC<VideoChatProps> = ({ wsUrl, userId }) => {
  const localVideoRef = useRef<HTMLVideoElement>(null);
  const wsRef = useRef<WebSocket | null>(null);
  const [peers, setPeers] = useState<PeerConnections>({});
  const videoRefs = useRef<VideoElements>({});

  useEffect(() => {
    const init = async () => {
      // 1. Получаем локальный медиа-поток
      const stream = await navigator.mediaDevices.getUserMedia({ video: true, audio: true });
      if (localVideoRef.current) {
        localVideoRef.current.srcObject = stream;
      }

      // 2. Подключаемся к WebSocket
      const ws = new WebSocket(wsUrl);
      wsRef.current = ws;

      ws.onopen = () => {
        ws.send(JSON.stringify({ type: "join", id: userId }));
      };

      ws.onmessage = async (event) => {
        const msg = JSON.parse(event.data);
        const { type, from, data } = msg;

        if (from === userId) return; // Игнорируем свои сообщения

        switch (type) {
          case "join":
            // Новый пользователь присоединился, создаем PeerConnection
            await createOffer(from, stream, ws);
            break;

          case "offer":
            await handleOffer(from, data, stream, ws);
            break;

          case "answer":
            await handleAnswer(from, data);
            break;

          case "candidate":
            await handleCandidate(from, data);
            break;

          default:
            break;
        }
      };
    };

    init();

    return () => {
      Object.values(peers).forEach((pc) => pc.close());
      wsRef.current?.close();
    };
  }, []);

  const createOffer = async (remoteId: string, stream: MediaStream, ws: WebSocket) => {
    const pc = new RTCPeerConnection();
    peers[remoteId] = pc;
    setPeers({ ...peers });

    // Добавляем локальные треки
    stream.getTracks().forEach((track) => pc.addTrack(track, stream));

    // Приём удалённого потока
    pc.ontrack = (event) => {
      let videoEl = videoRefs.current[remoteId];
      if (!videoEl) {
        videoEl = document.createElement("video");
        videoEl.autoplay = true;
        videoEl.playsInline = true;
        videoRefs.current[remoteId] = videoEl;
        document.getElementById("videos")?.appendChild(videoEl);
      }
      videoEl.srcObject = event.streams[0];
    };

    pc.onicecandidate = (event) => {
      if (event.candidate) {
        ws.send(JSON.stringify({ type: "candidate", from: userId, to: remoteId, data: event.candidate }));
      }
    };

    const offer = await pc.createOffer();
    await pc.setLocalDescription(offer);
    ws.send(JSON.stringify({ type: "offer", from: userId, to: remoteId, data: offer }));
  };

  const handleOffer = async (remoteId: string, offer: RTCSessionDescriptionInit, stream: MediaStream, ws: WebSocket) => {
    const pc = new RTCPeerConnection();
    peers[remoteId] = pc;
    setPeers({ ...peers });

    // Добавляем локальные треки
    stream.getTracks().forEach((track) => pc.addTrack(track, stream));

    // Приём удалённого потока
    pc.ontrack = (event) => {
      let videoEl = videoRefs.current[remoteId];
      if (!videoEl) {
        videoEl = document.createElement("video");
        videoEl.autoplay = true;
        videoEl.playsInline = true;
        videoRefs.current[remoteId] = videoEl;
        document.getElementById("videos")?.appendChild(videoEl);
      }
      videoEl.srcObject = event.streams[0];
    };

    pc.onicecandidate = (event) => {
      if (event.candidate) {
        ws.send(JSON.stringify({ type: "candidate", from: userId, to: remoteId, data: event.candidate }));
      }
    };

    await pc.setRemoteDescription(new RTCSessionDescription(offer));
    const answer = await pc.createAnswer();
    await pc.setLocalDescription(answer);
    ws.send(JSON.stringify({ type: "answer", from: userId, to: remoteId, data: answer }));
  };

  const handleAnswer = async (remoteId: string, answer: RTCSessionDescriptionInit) => {
    const pc = peers[remoteId];
    if (!pc) return;
    await pc.setRemoteDescription(new RTCSessionDescription(answer));
  };

  const handleCandidate = async (remoteId: string, candidate: RTCIceCandidateInit) => {
    const pc = peers[remoteId];
    if (!pc) return;
    await pc.addIceCandidate(new RTCIceCandidate(candidate));
  };

  return (
    <div>
      <h2>Видео чат</h2>
      <video ref={localVideoRef} autoPlay playsInline muted style={{ width: 200, height: 150, marginRight: 10 }} />
      <div id="videos" style={{ display: "flex", flexWrap: "wrap" }}></div>
    </div>
  );
};

export default VideoChat;
