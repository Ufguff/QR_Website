import { useEffect, useState } from "react";
import LoadingCircle from "./components/LoadingCircle/LoadingCircle.tsx";
import { API_URL } from "./environment.ts";
import axios from "axios";
import { QRCode } from "./types.ts";
import ColorPicker from "./components/ColorPicker/ColorPicker.tsx";

function App() {
  const [textToQR, setTextToQR] = useState<string>(location.toString())
  const [QRSize, setQRSize] = useState<number>(250)
  const [QRImageUrl, setQRImageUrl] = useState<string>("")
  const [isLoading, setIsLoading] = useState<boolean>(false)

  const [backgroundColor, setBackgroundColor] = useState("#ffffff")
  const [foregroundColor, setForegroundColor] = useState("#000000")


  const getQR = async (): Promise<void> => {
    if (!textToQR || QRSize < 250 || QRSize > 2048) return

    setIsLoading(true)
    const endpoint = "/api/v1/GetQR"
    const url = API_URL + endpoint
    const body: QRCode = {
      url: textToQR,
      size: QRSize,
      background: backgroundColor,
      foreground: foregroundColor
    }

    const response = await axios.post(url, body, {responseType: "blob"});
    const imageFile: File = new File([response.data], 'image.png', {type: "image/png"});
    setQRImageUrl(URL.createObjectURL(imageFile))

    setIsLoading(false)
  }

  useEffect(() => {
    getQR()
  }, [])


  return (
    <div id="app" className="full-size">
      <div className="qr-form">
        <div className="qr-form-wrapper">
          <div className="qr-form-left">
            <h2>QR-code creator</h2>
            <div className="qr-form-edit">
              <div className="qr-form-edit-section">
                <p>Введите URL-адрес</p>
                <input
                  type="text"
                  className="qr-form-text-input"
                  placeholder="https://"
                  value={textToQR}
                  onChange={e => setTextToQR(e.currentTarget.value)}
                />
              </div>
              <div className="qr-form-edit-section">
                <input
                  min={250}
                  max={2048}
                  className="slider"
                  type="range"
                  value={QRSize}
                  onChange={(e) => setQRSize(parseInt(e.currentTarget.value))}
                />
                <p className="text-center">{QRSize}x{QRSize} Px</p>
              </div>
              <div className="qr-form-edit-section">
                <div className="qr-form-edit-color">
                  <div>Цвет QR кода:</div>
                  <ColorPicker color={foregroundColor} setColor={setForegroundColor}/>
                </div>
                <div className="qr-form-edit-color">
                  <div>Цвет фона:</div>
                  <ColorPicker color={backgroundColor} setColor={setBackgroundColor}/>
                </div>
              </div>
            </div>
            <button id="create-button" className="btn" onClick={getQR}>Создать QR-Код</button>
          </div>
          <div className="qr-form-answer">
            <div className="qr-form-answer-image content-center">
              <div className="qr-form-answer-image-bg">
                {isLoading && (<LoadingCircle/>)}
                {!isLoading && (<img src={QRImageUrl} alt=""/>)}
              </div>
            </div>
            <div className="qr-form-answer-buttons">
              <a href={QRImageUrl} download={"qr_code.png"}>
                <button className="btn download-button">Скачать PNG</button>
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
