// use digital_outlet

// create
// a
db.operators.insertMany([
  {
    _id: 1,
    name: "operator satu",
    created_at: new Date(),
  },
  {
    _id: 2,
    name: "operator dua",
    created_at: new Date(),
  },
  {
    _id: 3,
    name: "operator tiga",
    created_at: new Date(),
  },
  {
    _id: 4,
    name: "operator empat",
    created_at: new Date(),
  },
  {
    _id: 5,
    name: "operator lima",
    created_at: new Date(),
  },
]);

// b
db.product_types.insertMany([
  {
    _id: 1,
    name: "product type satu",
    created_at: new Date(),
  },
  {
    _id: 2,
    name: "product type dua",
    created_at: new Date(),
  },
  {
    _id: 3,
    name: "product type tiga",
    created_at: new Date(),
  },
]);

// c
db.product.insertMany([
  {
    _id: 1,
    _product_types_id: 1,
    _operators_id: 3,
    code: Math.random(),
    status: 1,
    name: "product satu",
    created_at: new Date(),
  },
  {
    _id: 2,
    _product_types_id: 1,
    _operators_id: 3,
    code: Math.random(),
    status: 1,
    name: "product dua",
    created_at: new Date(),
  },
]);

// d
db.product.insertMany([
  {
    _id: 3,
    _product_types_id: 2,
    _operators_id: 1,
    code: Math.random(),
    status: 1,
    name: "product tiga",
    created_at: new Date(),
  },
  {
    _id: 4,
    _product_types_id: 2,
    _operators_id: 1,
    code: Math.random(),
    status: 1,
    name: "product empat",
    created_at: new Date(),
  },
  {
    _id: 5,
    _product_types_id: 2,
    _operators_id: 1,
    code: Math.random(),
    status: 1,
    name: "product lima",
    created_at: new Date(),
  },
]);

// e
db.product.insertMany([
  {
    _id: 6,
    _product_types_id: 3,
    _operators_id: 4,
    code: Math.random(),
    status: 1,
    name: "product enam",
    created_at: new Date(),
  },
  {
    _id: 7,
    _product_types_id: 3,
    _operators_id: 4,
    code: Math.random(),
    status: 1,
    name: "product tujuh",
    created_at: new Date(),
  },
  {
    _id: 8,
    _product_types_id: 3,
    _operators_id: 4,
    code: Math.random(),
    status: 1,
    name: "product delapan",
    created_at: new Date(),
  },
]);

// f
db.product.updateMany(
  {},
  {
    $set: { description: "this is description" },
  },
  {
    upsert: true,
  }
);

// g
db.payment_methods.insertMany([
  {
    _id: 1,
    name: "payment method satu",
    status: 1,
    created_at: new Date(),
  },
  {
    _id: 2,
    name: "payment method dua",
    status: 1,
    created_at: new Date(),
  },
  {
    _id: 3,
    name: "payment method tiga",
    status: 1,
    created_at: new Date(),
  },
]);

// h
db.users.insertMany([
  {
    _id: 1,
    name: "user satu",
    status: 1,
    dob: new Date(),
    gender: "laki-laki",
    created_at: new Date(),
  },
  {
    _id: 2,
    name: "user dua",
    status: 1,
    dob: new Date(),
    gender: "laki-laki",
    created_at: new Date(),
  },
  {
    _id: 3,
    status: 1,
    name: "user tiga",
    dob: new Date(),
    gender: "laki-laki",
    created_at: new Date(),
  },
  {
    _id: 4,
    name: "user empat",
    status: 1,
    dob: new Date(),
    gender: "perempuan",
    created_at: new Date(),
  },
  {
    _id: 5,
    name: "user lima",
    status: 1,
    dob: new Date(),
    gender: "perempuan",
    created_at: new Date(),
  },
]);

// i dan j
// db.users.updateMany(
//   {},
//   {
//     $set: {
//       transactions: [
//         {
//           _payment_methods_id: 1,
//           status: 1,
//           total_qty: 3,
//           total_price: 2000000,
//           created_at: new Date(),
//           products: [
//             {
//               product_type: 1,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 1",
//               status: 1,
//               created_at: new Date(),
//             },
//             {
//               product_type: 2,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 2",
//               status: 1,
//               created_at: new Date(),
//             },
//             {
//               product_type: 1,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 3",
//               status: 1,
//               created_at: new Date(),
//             },
//           ],
//         },
//         {
//           _payment_methods_id: 2,
//           status: 0,
//           total_qty: 3,
//           total_price: 3500000,
//           created_at: new Date(),
//           products: [
//             {
//               product_type: 1,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 1",
//               status: 1,
//               created_at: new Date(),
//             },
//             {
//               product_type: 2,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 2",
//               status: 1,
//               created_at: new Date(),
//             },
//             {
//               product_type: 1,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 3",
//               status: 1,
//               created_at: new Date(),
//             },
//           ],
//         },
//         {
//           _payment_methods_id: 1,
//           status: 1,
//           total_qty: 3,
//           total_price: 1500000,
//           created_at: new Date(),
//           products: [
//             {
//               product_type: 1,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 1",
//               status: 1,
//               created_at: new Date(),
//             },
//             {
//               product_type: 2,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 2",
//               status: 1,
//               created_at: new Date(),
//             },
//             {
//               product_type: 1,
//               operator_id: 1,
//               code: Math.random().toString(15),
//               name: "product 3",
//               status: 1,
//               created_at: new Date(),
//             },
//           ],
//         },
//       ],
//     },
//   },
//   { upsert: true }
// );

db.transactions.insertMany([
  {
    _user_id: 1,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 1,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 1,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 3,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 3,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 3,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 2,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 3,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 3,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 3,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 2,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 2,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 2,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 2,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 2,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 4,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 4,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 4,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 5,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 5,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
  {
    _user_id: 5,
    _payment_methods_id: 1,
    status: 1,
    total_qty: 2,
    total_price: 2000000,
    created_at: new Date(),
    products: [
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 1",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 2",
        status: 1,
        created_at: new Date(),
      },
      {
        _id: 1,
        product_type: 1,
        operator_id: 1,
        code: Math.random().toString(15),
        name: "product 3",
        status: 1,
        created_at: new Date(),
      },
    ],
  },
]);

// read
// a
db.users.find({ gender: "laki-laki" });

// b
db.product.find({ _id: 3 });

// c
db.users.countDocuments({ gender: "perempuan" });

// e
db.users.find().sort({ name: 1 }); //asc=1 desc=-1

// f
db.product.find().limit(5); //asc=1 desc=-1

// update
// a
db.product.updateOne({ _id: 1 }, { $set: { name: "product dummy" } }, {});

// b
db.transactions.updateMany(
  { "products._id": 1 },
  { $set: { total_qty: 3 } },
  {}
);

// delete
// a
db.product.deleteMany({
  _id: 1,
});
// b
db.product.deleteMany({
  _product_types_id: 1,
});
