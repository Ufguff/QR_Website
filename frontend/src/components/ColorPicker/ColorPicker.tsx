import "./ColorPicker.scss"
import { HexColorPicker } from "react-colorful";
import React, { FC, useRef, useState } from "react";
import { useOutsideClick } from "../../hooks/useOutsideClick.ts";

interface ColorPickerProps {
  color: string;
  setColor: React.Dispatch<React.SetStateAction<string>>;
}

const ColorPicker: FC<ColorPickerProps> = ({color, setColor}) => {
  const [active, setActive] = useState<boolean>(false)
  const colorPickerRef = useRef(null);
  useOutsideClick(colorPickerRef, () => setActive(false))

  return (
    <div className="color-picker" onClick={() => setActive(true)}>
      <div className="color-picker-wrapper" style={{backgroundColor: color}}>
        <div className="color-picker-overlay" style={{display: active ? "block" : "none"}}>
          <div ref={colorPickerRef}>
            <HexColorPicker color={color} onChange={setColor}/>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ColorPicker;