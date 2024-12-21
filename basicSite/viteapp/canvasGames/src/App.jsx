import { useEffect, useState, useRef } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function randint(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function App() {
  const canvasRef = useRef();
  const [snakeList, setSnakeList] = useState([{x: 100, y: 110}]);
  const [foodPos, setFoodPos] = useState({x: randint(20, 980), y: randint(20, 680)});
  const [direction, setDirection] = useState('right');
  const [score, setScore] = useState(0);

  const nodeLength = 15;
  const nodeHeight = 15;
  
  useEffect(() => {
    const gameLoop = setInterval(() => {
      const canvas = canvasRef.current;
      const ctx = canvas.getContext('2d');

      // canvas dimensions
      canvas.width = 1000;
      canvas.height = 700;

      // canvas background
      ctx.fillStyle = 'black';
      ctx.fillRect(0, 0, canvas.width, canvas.height);

      // do snake game stuff

      // draw snake
      ctx.beginPath();
      ctx.rect(30, 40, nodeLength, nodeHeight);
      ctx.fillStyle = 'purple';
      ctx.fill();
      ctx.closePath();

      // draw food
      ctx.beginPath();
      ctx.rect(foodPos.x, foodPos.y, nodeLength, nodeHeight);    
      ctx.fillStyle = 'red';
      ctx.fill();
      ctx.closePath();
      }, 1000 / 60);
  }, [snakeList, direction, foodPos, score]);

  return (
    <div id="canvasContainer">
      <canvas ref={canvasRef} />
    </div>
  )
}

export default App
