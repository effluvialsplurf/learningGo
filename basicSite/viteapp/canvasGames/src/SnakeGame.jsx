import { useEffect, useState, useRef } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

// up top here we have some helper functions
function randint(min, max) {
  return Math.floor(Math.random() * (max - min + 1)) + min;
}

function normalizePos(pos) {
  return pos - (pos % 15);
}

// main component, a snake game
export default function SnakeGame() {
  // initializing our state variables
  const canvasRef = useRef();
  const [snakeList, setSnakeList] = useState([{x: 120, y: 150}]);
  const [foodPos, setFoodPos] = useState({x: normalizePos(randint(20, 980)), y: normalizePos(randint(20, 680))});
  const [direction, setDirection] = useState('right');
  const [score, setScore] = useState(0);

  // we use increments of 15 for our sizing
  const nodeLength = 15;
  const nodeHeight = 15;

  // initialize the game
  const initializeGame = () => {
    setSnakeList([{x: 120, y: 150}]);
    setFoodPos({x: normalizePos(randint(20, 980)), y: normalizePos(randint(20, 680))});
    setDirection('right');
    setScore(0);
  }

  // draw the canvas
  const drawCanvas = () => {
    const canvas = canvasRef.current;
    const ctx = canvas.getContext('2d');

    // canvas dimensions
    canvas.width = 1000;
    canvas.height = 700;

    // canvas background
    ctx.fillStyle = 'black';
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    // draw snake
    for (let i = 0; i < snakeList.length; i++) {
      ctx.beginPath();
      ctx.rect(snakeList[i].x, snakeList[i].y, nodeLength, nodeHeight);
      ctx.fillStyle = 'purple';
      ctx.fill();
      ctx.closePath();
    }

    // draw food
    ctx.beginPath();
    ctx.rect(foodPos.x, foodPos.y, nodeLength, nodeHeight);    
    ctx.fillStyle = 'red';
    ctx.fill();
    ctx.closePath();
  }

  // update snake direction and returns an altered node
  const updateSnakeDirection = (snakeNode, direction) => {
    if (direction === 'up') {
      snakeNode.y -= 15;
    } else if (direction === 'down') {
      snakeNode.y += 15;
    } else if (direction === 'left') {
      snakeNode.x -= 15;
    } else if (direction === 'right') {
      snakeNode.x += 15;
    }
    
    return snakeNode;
  }

  // adds a new node to the snake list
  // has to do some logic to determine where to add the node
  const addSnakeNode = (snakeList) => {
    let list = [...snakeList];
    let alterx = 0;
    let altery = 0;

    if (direction === 'up') {
      altery = 15;
    } else if (direction === 'down') {
      altery = -15;
    } else if (direction === 'left') {
      alterx = 15;
    } else if (direction === 'right') {
      alterx = -15;
    }

    list.push({x: snakeList[snakeList.length - 1].x + alterx, y: snakeList[snakeList.length - 1].y + altery});
    setSnakeList([...list]);
  }

  // check for collision, pretty much all our 'game' logic is here
  const collision = () => {
    if (snakeList[0].x === foodPos.x && snakeList[0].y === foodPos.y) {
      setScore(score + 1);
      setFoodPos({x: normalizePos(randint(20, 980)), y: normalizePos(randint(20, 680))});
      addSnakeNode(snakeList);
    }

    if (snakeList[0].x < 0 || snakeList[0].x > 1000 || snakeList[0].y < 0 || snakeList[0].y > 700) {
      initializeGame();
    }

    for (let i = 1; i < snakeList.length; i++) {
      if (snakeList[0].x === snakeList[i].x && snakeList[0].y === snakeList[i].y) {
        initializeGame();
        break;
      }
    }
  }

  // this handles any keypresses, adds an event listener on mount
  useEffect (() => {
    const handleKeyDown = (e) => {
      switch (e.key) {
        case 'ArrowUp':
          setDirection('up');
          break;
        case 'ArrowDown':
          setDirection('down');
          break;
        case 'ArrowLeft':
          setDirection('left');
          break;
        case 'ArrowRight':
          setDirection('right');
          break;
      }
    };

    window.addEventListener('keydown', handleKeyDown);

    return () => {
      window.removeEventListener('keydown', handleKeyDown);
    };
  }, []);
  
  // the main game loop set at 20 fps
  useEffect(() => {
    const gameLoop = setInterval(() => {
      // update snake position, the direction of the head and set all other positions to prev pos of next up
      let newSnakeList = [...snakeList];
      for (let i = snakeList.length - 1; i >= 0; i--) {
        if (i === 0) {
          updateSnakeDirection(newSnakeList[i], direction);
        } else {
          newSnakeList[i].x = snakeList[i - 1].x;
          newSnakeList[i].y = snakeList[i - 1].y;
        }
      }
      setSnakeList([...snakeList]);

      // check for collision
      collision();

      // render the game
      drawCanvas();
    }, 1000 / 20);

    return () => {
      clearInterval(gameLoop);
    };
  }, [direction, score, snakeList, foodPos]);

  return (
    <>
      <div id="scoreContainer">Score: {score}</div>
      <div id="canvasContainer">
        <canvas ref={canvasRef} />
      </div>
    </>
  )
}