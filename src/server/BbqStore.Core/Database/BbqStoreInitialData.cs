using System;
using System.Collections.Generic;
using System.Text;
using BbqStore.Core.Entities;
using Marten;
using Marten.Schema;

namespace BbqStore.Core.Database
{
    public class BbqStoreInitialData : IInitialData
    {
        private object[] InitialData { get; set; }

        public BbqStoreInitialData(params object[] initialData)
        {
            InitialData = initialData;
        }

        public void Populate(IDocumentStore store)
        {
            using (var session = store.LightweightSession())
            {
                session.Store(InitialData);
                session.SaveChanges();
            }
        }
    }

    public static class InitialDataSets
    {
        public static readonly Product[] Products =
        {
            new Product
            {
                Id = Guid.Parse("00000000-0000-0000-0000-000000000001"),
                Name = "Beef Brisket",
                Key = "beef-brisket",
                Description = "Texas Style Prime Brisket.  Cooked low and slow for 12+ hours",
                Price = 21.50m,
                Unit = "Pound",
                CreatedBy = "chef",
                ModifiedBy = "chef",
                CreatedDate = DateTimeOffset.Now,
                ModifiedDate = DateTimeOffset.Now
            },
            new Product
            {
                Id = Guid.Parse("00000000-0000-0000-0000-000000000002"),
                Name = "Pulled Pork",
                Key = "pulled-pork",
                Description = "North Carolina Style Pulled Pork.  Cooked low and slow for 12+ hours",
                Price = 16.50m,
                Unit = "Pound",
                CreatedBy = "chef",
                ModifiedBy = "chef",
                CreatedDate = DateTimeOffset.Now,
                ModifiedDate = DateTimeOffset.Now
            },
            new Product
            {
                Id = Guid.Parse("00000000-0000-0000-0000-000000000003"),
                Name = "Whole Chicken",
                Key = "whole-chicken",
                Description = "Smoked Chicken.  Cooked low and slow for 4+ hours",
                Price = 17.25m,
                Unit = "Whole Chicken",
                CreatedBy = "chef",
                ModifiedBy = "chef",
                CreatedDate = DateTimeOffset.Now,
                ModifiedDate = DateTimeOffset.Now
            },
            new Product
            {
                Id = Guid.Parse("00000000-0000-0000-0000-000000000004"),
                Name = "Half Chicken",
                Key = "half-chicken",
                Description = "Smoked Chicken.  Cooked low and slow for 4+ hours",
                Price = 9.75m,
                Unit = "Half Chicken",
                CreatedBy = "chef",
                ModifiedBy = "chef",
                CreatedDate = DateTimeOffset.Now,
                ModifiedDate = DateTimeOffset.Now
            }
        };

        public static readonly Store[] Stores =
        {
            new Store
            {
                Id = Guid.Parse("00000000-0000-0000-0000-000000000001"),
                Name = "BBQ 52",
                Key = "bbq-52",
                CreatedBy = "chef",
                ModifiedBy = "chef",
                CreatedDate = DateTimeOffset.Now,
                ModifiedDate = DateTimeOffset.Now
            }
        };
    }
}