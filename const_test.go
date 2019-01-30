package xlripper

const strRowCloseTest = `<c r="A2" s="1" t="s"><v>0</v></c><c r="B2" s="1" t="s"><v>0</v></c><c r="C2" s="1" t="s"><v>0</v></c><c r="D2" s="1" t="s"><v>2</v></c><c r="E2" s="1" t="s"><v>1</v></c><c r="F2" s="1" t="s"><v>1</v></c><c r="G2" s="1" t="s"><v>1</v></c><c r="H2" s="1" t="s"><v>3</v></c><c r="I2" s="1" t="s"><v>2</v></c><c r="J2" t="str"><f t="shared" si="1"/><v>it('should identify column type NULL NULL NULL STRING', async (done) =&gt; {
  const data = new Map();
  data.set(0, { col: null });
  data.set(1, { col: null });
  data.set(2, { col: null });
  data.set(3, { col: &quot;X&quot; });
  const table = new Table();
  table.idMap = data;
  table.updateHeaderMetadata();
  const header = table.headers.getByName('col');
  expect(header).toBeTruthy();
  expect(header).toHaveProperty('name', 'col');
  expect(header).toHaveProperty('type', VALUE_TYPES.STRING);
  expect(header).toHaveProperty('index', 0); done();
})</v></c></row><row r="3"><c r="A3" s="1" t="s"><v>0</v></c><c r="B3" s="1" t="s"><v>0</v></c><c r="C3" s="1" t="s"><v>0</v></c><c r="D3" s="1" t="s"><v>4</v></c><c r="E3" s="1" t="s"><v>1</v></c><c r="F3" s="1" t="s"><v>1</v></c><c r="G3" s="1" t="s"><v>1</v></c><c r="H3" s="2"><v>2.1</v></c><c r="I3" s="1" t="s"><v>4</v></c><c r="J3" t="str"><f t="shared" si="1"/><v>it('should identify column type NULL NULL NULL DECIMAL', async (done) =&gt; {
  const data = new Map();
  data.set(0, { col: null });
  data.set(1, { col: null });
  data.set(2, { col: null });
  data.set(3, { col: 2.1 });
  const table = new Table();`
