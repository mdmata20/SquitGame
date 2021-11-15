var express = require('express');
var router = express.Router();

const gamesController = require('../controllers/controladores')

router.get('/api/games/count', gamesController.GetAllGamesCount);
router.get('/api/games', gamesController.GetAllGames);
router.get('/api/games/top10', gamesController.GetLast10Games);
router.get('/api/games/top10players', gamesController.GetBest10Players);
router.get('/api/games/top3games', gamesController.GetTop3Games);
router.get('/api/games/workers', gamesController.GetWorkers);
router.get('/api/games/:winner', gamesController.GetAllGamesByWinner);

module.exports = router;
