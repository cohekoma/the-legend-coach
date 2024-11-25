# Winning Probability

The core of this game will be based on the probability theory. Mainly, it will be about the possibility of winning or the winning rate between two teams in a match. Besides this probability, there will be many different probability, for example: the probability of a player to be a star (or his potential), the probability of a passing (can it be successfully or not?), the probability of a shooting (can it score a goal?), etc.

But within the scope of this page, I will try to do the research on how I can implement the winning probability of a team within a match.

## The basic

To begin with this, I have to say first that I don't know at this point that how to implement this. However, the very initial thought and idea that I came up with would be like this:

Say today, right at the time I'm writing this (24th November, 2024), we are about to have a match between Manchester United vs Ipswich Town. Right away, my thought for this would be that the winning possibility of Manchester United would definitely be higher, meaning that their winning rate must be higher than Ipswich Town. If I have to put this into a number, it might be 90% winning rate for Manchester United and only 10% of winning rate for Ipswich Town.

How did I come up with that number and the possibility of winning rate that is very leaning towards Manchester United?

### Football Language

In order to answer this question, I need to first get rid of all the other factors that might cause the unpredictability. Let me jot it down first all the aspects that I think are not the unpredictable factors in the football language:
- Chart Position: Man Utd is at 12th position, while Ipswich Town is 19th - the second last position (20th would be the bottom of the table).
- Recent Stats: Man Utd has 4W(win)-4D(draw)-4L(lost) after 11 games, while Ipswich Town has 1W-5D-5L after 11 games.

The factors above are based on numbers, statistics, which make them certain and can support the winning possibility of Manchester United. Now, let's think of the unpredictable factors:
- Morale: Players' morale can definitely be unpredictable. Right at this moment, you can think that Man Utd has a higher morale in general because they just hired a new coach (Ruben Amorim) and the atmostphere of the club is very good at the moment. However, I can say the same thing about Ipswich Town's morale. They have just been promoted this year and even when they had only won 1 game out of 11 games, they still have every chance to stay at the league. Furthermore, losing too many games can make the players feel like they have to win, which can *boost* their engines. You can make any similar argument about this based on the context, so it's very unpredictable.
- Playing Form: Same with morality, playing form is definitely not predictable. Good players can have bad days while bad players can have excellent days. Today might be that day.
- Formation and Strategy: This is super unpredictable and these aspects definitely will affect the player's forms and performances. Because I'm a Manchester United's fan so I can give you an example about them. Ruben Amorim, the new head coach is very famous for his 3-4-3 formation while Manchester United is very famous for 4-4-2 fomartion. The current players of Man Utd don't seem to be accustomed to the formation with 3 players in the back, so it might be difficult for them.

### Implementation

What I write above are very football-ish. We are trying to implement a game so we must try and convert everything into numbers. Let's get back to the 90% winning rate for Man Utd and 10% winning rate for Ipswich Town that I give you above. When I write this, I don't have any stats supported me yet. This is based pretty much on assumption. In order to implement this assumption into our game, we definitely need to make our thoughts *"numeric"*.

In games like FIFA, we shall have something like an overall score / stat / point assigned to each team, which indicates how strong a team is. So, I think at the moment, I will try to assign this stat for each team. 

### Update on 25th Nov

One point that I missed when I was writing about the algorithm to find the possibility was that I completely forgot about the drawing / tied result. Which then led to one interesting result: Man Utd didn't win Ipswich Town, the result was a tied (1-1).