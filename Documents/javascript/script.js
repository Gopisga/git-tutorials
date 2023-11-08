/*var number = 56;
var number1 = 67;
console.log(number + number1);

 var str1 = "thisis a string";
 var str2 = "thisd me ";
 console.log(str1+str2);

 var marks = {
    gopi :1,
    arpan :2,
    ankit :3,
    prashant :4
 }
 console.log(marks)

 var arr = [1,2,3,4,5,6]
  console.log(arr[2+2])

  var age =4
  if (age > 18){
  console.log('you can see h');
  }else{
    console.log('you cant see h')
  }*/

  class TwitterX {
    constructor(name) {
      this.name = name;
      this.tweets = [];
    }
  
    // Method to add a new tweet
    addTweet(tweet) {
      this.tweets.push(tweet);
    }
  
    // Method to get all tweets
    getAllTweets() {
      return this.tweets;
    }
  }
  
  // Create an instance of TwitterX
  const twitterX = new TwitterX("TwitterX");
  
  // Add tweets to the instance
  twitterX.addTweet("First tweet");
  twitterX.addTweet("Second tweet");
  twitterX.addTweet("Third tweet");
  
  // Get all tweets
  const allTweets = twitterX.getAllTweets();
  console.log(allTweets);
