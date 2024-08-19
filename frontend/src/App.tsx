import { useState } from "react";

function App() {
  const [textToQR, setTextToQR] = useState<string>("")

  return (
    <div id="app" className="full-size">
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
            <div className="qr-form-answer-image content-center">
              <div className="qr-form-answer-image-bg">
                <img src="" alt=""/>
              </div>
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
