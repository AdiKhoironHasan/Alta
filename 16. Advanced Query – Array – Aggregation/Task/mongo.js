db.books.insertMany([
  {
    _id: 1,
    title: "Wawasan Pancasila",
    authoriD: 1,
    publisherID: 1,
    price: 71200,
    stats: { page: 324, weight: 300 },
    publishedAt: new Date("2018-10-01"),
    category: ["social", "politics"],
  },
  {
    _id: 2,
    title: "Mate Air Keteladanan",
    authoriD: 1,
    publisheriD: 2,
    price: 106250,
    stats: { page: 672, weight: 650 },
    publishedAt: new Date("2017-09-01"),
    category: ["social", "politics"],
  },
  {
    _id: 3,
    title: "Revolusi Pancasila",
    authoriD: 1,
    publisheriD: 1,
    price: 54400,
    stats: { page: 220, weight: 500 },
    publishedAt: new Date("2015-05-01"),
    category: ["social", "politics"],
  },
  {
    _id: 4,
    title: "Self Driving",
    authoriD: 2,
    publisheriD: 1,
    price: 58650,
    stats: { page: 286, weight: 300 },
    publishedAt: new Date("2018-05-01"),
    category: ["self-development"],
  },
  {
    _id: 5,
    title: "Self Disruption",
    authoriD: 2,
    publisheriD: 2,
    price: 83300,
    stats: { page: 400, weight: 800 },
    publishedAt: new Date("2018-05-01"),
    category: ["self-development"],
  },
]);

db.authors.insertMany([
  { _id: 1, firstName: "Yudi", lastName: "Latif" },
  { _id: 2, firstName: "Rhenald", lastName: "Kasali" },
]);

db.publishers.insertMany([
  { _id: 1, publisherName: "Expose" },
  { _id: 2, publisherName: "Mizan" },
]);

// 1
db.books.find({ authoriD: 2 });

// 2
db.books.find(
  {
    authoriD: 1,
  },
  {
    _id: 1,
    title: 1,
    price: 1,
  }
);

// 3
db.books.aggregate([
  {
    $match: {
      authoriD: 2,
    },
  },
  {
    $group: {
      _id: null,
      total: { $sum: "$stats.page" },
    },
  },
]);

// 4
db.authors.aggregate([
  {
    $lookup: {
      from: "books",
      localField: "_id",
      foreignField: "authoriD",
      as: "allBooks",
    },
  },
]);

// 5
db.books.aggregate([
  {
    $lookup: {
      from: "authors",
      localField: "authoriD",
      foreignField: "_id",
      as: "authorData",
    },
  },
  { $unwind: "$authors" },
  {
    $lookup: {
      from: "publishers",
      localField: "publisherID",
      foreignField: "_id",
      as: "publisherData",
    },
  },
  { $unwind: "$publishers" },
]);

// 6
db.books.aggregate([
  {
    $lookup: {
      from: "authors",
      localField: "authoriD",
      foreignField: "_id",
      as: "authorData",
    },
  },
  {
    $group: {
      _id: "$authorData.firstName",
      number_of_books: { $sum: 1 },
    },
  },
]);

db.books.aggregate([
  {
    $lookup: {
      from: "authors",
      localField: "authoriD",
      foreignField: "_id",
      as: "authorData",
    },
  },
  {
    $project: {
      number_of_books: { $sum: 1 },
      title: "$title",
      price: "$price",
      discount: {
        $cond: {
          if: {
            $gt: ["price", 90000],
          },
          then: 30,
          else: {
            if: {
              $gt: ["price", 60000],
            },
            then: 20,
            else: 10,
          },
        },
      },
    },
  },
]);

// 7
db.books.aggregate([
  {
    $lookup: {
      from: "authors",
      localField: "authoriD",
      foreignField: "_id",
      as: "authorData",
    },
  },
  {
    $project: {
      number_of_books: { $sum: 1 },
      title: "$title",
      price: "$price",
      discount: {
        $switch: {
          branches: [
            { case: { $gt: ["$price", 90000] }, then: "3%" },
            { case: { $gt: ["$price", 60000] }, then: "2%" },
            { case: { $gt: ["$price", 0] }, then: "1%" },
          ],
          default: 0,
        },
      },
    },
  },
]);

// 8
db.books.aggregate([
  {
    $lookup: {
      from: "authors",
      localField: "authoriD",
      foreignField: "_id",
      as: "authorData",
    },
  },
  {
    $lookup: {
      from: "publishers",
      localField: "publisherID",
      foreignField: "_id",
      as: "publisherData",
    },
  },
  {
    $project: {
      _id: null,
      title: "$title",
      price: "$price",
      author: "$authorData.firstName",
      publisher: "$publisherData.publisherName",
    },
  },
]);

// 9
db.books.aggregate([
  {
    $lookup: {
      from: "publishers",
      localField: "publisherID",
      foreignField: "_id",
      as: "publisherData",
    },
  },
  // {$unwind: "$publisherData"},
  {
    $match: {
      _id: {
        $in: [3, 4],
      },
    },
  },
  {
    $project: {
      title: "$title",
      price: "$price",
      publisher: "$publisherData.publisherName",
    },
  },
]);
