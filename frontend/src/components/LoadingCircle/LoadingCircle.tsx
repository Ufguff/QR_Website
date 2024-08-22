import { FC } from 'react';
import "./LoadingCircle.scss"

const LoadingCircle: FC = () => {
  return (
    <svg className="spinner" viewBox="0 0 50 50">
      <circle className="path" cx="25" cy="25" r="20" fill="none" stroke-width="8"></circle>
    </svg>
  );
};

export default LoadingCircle