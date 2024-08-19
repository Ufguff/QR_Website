import { useState } from "react";

function App() {
  const [textToQR, setTextToQR] = useState<string>("")

  return (
    <div id="app" className="full-size content-center">
      <div className="qr-form">
        <div className="qr-form-wrapper">
          <div className="qr-form-edit">
            <input
              type="text"
              className="qr-form-text-input"
              placeholder="URL"
              value={textToQR}
              onChange={e => setTextToQR(e.currentTarget.value)}
            />
            <button id="create-button" className="btn">Создать QR-Код</button>
          </div>
          <div className="qr-form-answer">
            <div className="qr-form-answer-image">
              <img src="" alt="" width="250" height="250"/>
            </div>
            <div className="qr-form-answer-buttons">
              <button className="btn download-button">Скачать PNG</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
