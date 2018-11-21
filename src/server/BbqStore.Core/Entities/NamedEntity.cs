namespace BbqStore.Core.Entities
{
    public class NamedEntity : Entity
    {
        public virtual string Key { get; set; }
        public virtual string Name { get; set; }
    }
}