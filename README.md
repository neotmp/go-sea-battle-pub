Sea Battle

The game logic follows conventional rules of the classical game:
Two players, each have five ships to deploy
Ships of different length: (5, 4, 3, 3, 2)
Ships can be placed either along an x- or y-axis.
The theater of war measures 10 by 10 cells.
Each player can see only his ships. The opponent’s ships are hidden unless they’re hit (direct hit is registered without exposing the ship until the ship is sunk).
Each player places their ships onto their designated grid and target that of their enemy.
Players exchange shots in turns by targeting specific cell of opposing grid.
Human player fires a first shot.
Every shot is either a hit or a miss.
Each ship can sustain maximum hits that correspond to their length. For instance, a destroyer that takes two hits is sunk.
The first player who loses their entire fleet loses the game.

I envision the game in three distinct modes

Simple mode allows human and computer exchange shots in turn. Computer deploys its ships randomly and fires its shots the same way.

Simple Mode with salvo. Instead of firing only one shot each side may fire up to five shots simultaneously. You can fire as many shots as number of ships you have. You lose a ship you lose one shot. The same logic applies to computer player. Besides that, there are no changes in how the game is played.

Advance Mode

Player’s logic remains the same. What’s different is how computer makes its decisions what cell to target. In Simple Mode it’s done randomly.

Now computer pursues a different strategy: Search and Destroy. At first it splits the area into smaller quadrants that it targets randomly. Once it hits a ship it switches to Destroy Mode, which means that it will keep targeting adjacent cells until the target is destroyed. So, in essence, it mimics human behavior.

I think it’s a really nice challenge from algorithm standpoint.

I’m not sure if I will implement salvo in Advance Mode. I’ll circle back to it later.

Statistical Mode (a.k.a ML, a.k.a AI)

Again, it is all about optimizing and enhancing computer’s ability to beat a human player. This time changes affect not only how computer target human’s fleet (offensive side), but also how computer deploy its capabilities (defensive side). Computer uses past data to find places on the grid that are least likely to be targeted by human. The same goes with Search and Destroy tactics. Computer targets areas that according to collected data have highest probabilities of having human fleet deployed there. Visually, it may be presented as a heatmap. You target the hottest areas and if you get some you go into a “Terminator Mode” and destroy enemy’s ship.

Apart from Human v. Computer Mode I also have plans to allow Human v. Human Mode. Where players can send invitations and meet one another to duke it out online.

That’s about it when it comes to backend. I’ll also have statistics to show to players afterwards. You know, the whole shebang.
